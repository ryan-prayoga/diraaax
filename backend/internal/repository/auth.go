package repository

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"

	"github.com/ryanprayoga/diraaax/backend/internal/domain"
	"github.com/ryanprayoga/diraaax/backend/internal/utils"
)

type AuthRepository struct {
	pool          *pgxpool.Pool
	sessionSecret string
}

func NewAuthRepository(pool *pgxpool.Pool, sessionSecret string) *AuthRepository {
	return &AuthRepository{
		pool:          pool,
		sessionSecret: sessionSecret,
	}
}

func (r *AuthRepository) VerifyPINAndCreateSession(
	ctx context.Context,
	pin string,
	token string,
	userAgent *string,
	ipAddress *string,
	expiresAt time.Time,
) (*domain.AuthSession, error) {
	tx, err := r.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	trimmedPIN := strings.TrimSpace(pin)

	var accessCode domain.AccessCode
	err = tx.QueryRow(ctx, `
		SELECT id, label, created_by, code_type, code_hint
		FROM access_codes
		WHERE is_active = true
		  AND code_type = 'pin'
		  AND (expires_at IS NULL OR expires_at > NOW())
		  AND (max_uses IS NULL OR used_count < max_uses)
		  AND crypt($1, code_hash) = code_hash
		ORDER BY id DESC
		LIMIT 1
	`, trimmedPIN).Scan(
		&accessCode.ID,
		&accessCode.Label,
		&accessCode.CreatedBy,
		&accessCode.CodeType,
		&accessCode.CodeHint,
	)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return nil, err
		}

		// Fallback for bcrypt-hashed PINs.
		accessCodePtr, bcryptErr := r.findActiveAccessCodeByPINBcrypt(ctx, tx, trimmedPIN)
		if bcryptErr != nil {
			return nil, bcryptErr
		}
		if accessCodePtr == nil {
			return nil, pgx.ErrNoRows
		}
		accessCode = *accessCodePtr
	}

	resolvedUserID, err := r.resolveSessionUserID(ctx, tx, accessCode.CreatedBy)
	if err != nil {
		return nil, err
	}
	accessCode.CreatedBy = resolvedUserID

	if _, err := tx.Exec(ctx, `
		UPDATE access_codes
		SET used_count = used_count + 1,
		    updated_at = NOW()
		WHERE id = $1
	`, accessCode.ID); err != nil {
		return nil, err
	}

	var session domain.Session
	err = tx.QueryRow(ctx, `
		INSERT INTO sessions (
			user_id,
			access_code_id,
			token_hash,
			user_agent,
			ip_address,
			expires_at,
			last_seen_at
		)
		VALUES ($1, $2, $3, $4, $5::inet, $6, NOW())
		RETURNING id, user_id, access_code_id, expires_at, last_seen_at, created_at
	`,
		resolvedUserID,
		accessCode.ID,
		utils.HashSessionToken(token, r.sessionSecret),
		userAgent,
		ipAddress,
		expiresAt,
	).Scan(
		&session.ID,
		&session.UserID,
		&session.AccessCodeID,
		&session.ExpiresAt,
		&session.LastSeenAt,
		&session.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	var user *domain.User
	if resolvedUserID != nil {
		user, err = r.getUserByID(ctx, tx, *resolvedUserID)
		if err != nil && !errors.Is(err, pgx.ErrNoRows) {
			return nil, err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &domain.AuthSession{
		Session:    &session,
		User:       user,
		AccessCode: &accessCode,
	}, nil
}

func (r *AuthRepository) resolveSessionUserID(ctx context.Context, tx pgx.Tx, accessCodeCreatedBy *int64) (*int64, error) {
	if accessCodeCreatedBy != nil {
		return accessCodeCreatedBy, nil
	}

	var userID int64
	err := tx.QueryRow(ctx, `
		SELECT id
		FROM users
		WHERE is_active = true
		ORDER BY id ASC
		LIMIT 1
	`).Scan(&userID)
	if err != nil {
		return nil, err
	}

	return &userID, nil
}

func (r *AuthRepository) findActiveAccessCodeByPINBcrypt(ctx context.Context, tx pgx.Tx, pin string) (*domain.AccessCode, error) {
	rows, err := tx.Query(ctx, `
		SELECT id, label, created_by, code_type, code_hint, code_hash
		FROM access_codes
		WHERE is_active = true
		  AND code_type = 'pin'
		  AND (expires_at IS NULL OR expires_at > NOW())
		  AND (max_uses IS NULL OR used_count < max_uses)
		ORDER BY id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			accessCode domain.AccessCode
			codeHash   string
		)

		if scanErr := rows.Scan(
			&accessCode.ID,
			&accessCode.Label,
			&accessCode.CreatedBy,
			&accessCode.CodeType,
			&accessCode.CodeHint,
			&codeHash,
		); scanErr != nil {
			return nil, scanErr
		}

		if !looksLikeBcryptHash(codeHash) {
			continue
		}

		if bcrypt.CompareHashAndPassword([]byte(codeHash), []byte(pin)) == nil {
			return &accessCode, nil
		}
	}

	if rowsErr := rows.Err(); rowsErr != nil {
		return nil, rowsErr
	}

	return nil, nil
}

func looksLikeBcryptHash(hash string) bool {
	return strings.HasPrefix(hash, "$2a$") ||
		strings.HasPrefix(hash, "$2b$") ||
		strings.HasPrefix(hash, "$2y$")
}

func (r *AuthRepository) GetSession(ctx context.Context, token string) (*domain.AuthSession, error) {
	var (
		session             domain.Session
		accessCodeID        *int64
		accessCodeLabel     *string
		accessCodeCreatedBy *int64
		accessCodeType      *string
		accessCodeHint      *string
		userID              *int64
		displayName         *string
		nickname            *string
		slug                *string
		avatarURL           *string
		isActive            *bool
		userCreatedAt       *time.Time
		userUpdatedAt       *time.Time
	)

	err := r.pool.QueryRow(ctx, `
		WITH active_session AS (
			UPDATE sessions
			SET last_seen_at = NOW()
			WHERE token_hash = $1
			  AND expires_at > NOW()
			RETURNING id, user_id, access_code_id, expires_at, last_seen_at, created_at
		)
		SELECT
			s.id,
			COALESCE(s.user_id, ac.created_by) AS actor_user_id,
			s.access_code_id,
			s.expires_at,
			s.last_seen_at,
			s.created_at,
			ac.id,
			ac.label,
			ac.created_by,
			ac.code_type,
			ac.code_hint,
			u.id,
			u.display_name,
			u.nickname,
			u.slug,
			u.avatar_url,
			u.is_active,
			u.created_at,
			u.updated_at
		FROM active_session s
		LEFT JOIN access_codes ac ON ac.id = s.access_code_id
		LEFT JOIN users u ON u.id = COALESCE(s.user_id, ac.created_by)
	`, utils.HashSessionToken(token, r.sessionSecret)).Scan(
		&session.ID,
		&session.UserID,
		&session.AccessCodeID,
		&session.ExpiresAt,
		&session.LastSeenAt,
		&session.CreatedAt,
		&accessCodeID,
		&accessCodeLabel,
		&accessCodeCreatedBy,
		&accessCodeType,
		&accessCodeHint,
		&userID,
		&displayName,
		&nickname,
		&slug,
		&avatarURL,
		&isActive,
		&userCreatedAt,
		&userUpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	var accessCode *domain.AccessCode
	if accessCodeID != nil && accessCodeLabel != nil && accessCodeType != nil {
		accessCode = &domain.AccessCode{
			ID:        *accessCodeID,
			Label:     *accessCodeLabel,
			CreatedBy: accessCodeCreatedBy,
			CodeType:  *accessCodeType,
			CodeHint:  accessCodeHint,
		}
	}

	var user *domain.User
	if userID != nil && displayName != nil && isActive != nil && userCreatedAt != nil && userUpdatedAt != nil {
		user = &domain.User{
			ID:          *userID,
			DisplayName: *displayName,
			Nickname:    nickname,
			Slug:        slug,
			AvatarURL:   avatarURL,
			IsActive:    *isActive,
			CreatedAt:   *userCreatedAt,
			UpdatedAt:   *userUpdatedAt,
		}
	}

	return &domain.AuthSession{
		Session:    &session,
		User:       user,
		AccessCode: accessCode,
	}, nil
}

func (r *AuthRepository) DeleteSession(ctx context.Context, token string) error {
	commandTag, err := r.pool.Exec(ctx, `
		DELETE FROM sessions
		WHERE token_hash = $1
	`, utils.HashSessionToken(token, r.sessionSecret))
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *AuthRepository) getUserByID(ctx context.Context, tx pgx.Tx, userID int64) (*domain.User, error) {
	var user domain.User
	err := tx.QueryRow(ctx, `
		SELECT id, display_name, nickname, slug, avatar_url, is_active, created_at, updated_at
		FROM users
		WHERE id = $1
	`, userID).Scan(
		&user.ID,
		&user.DisplayName,
		&user.Nickname,
		&user.Slug,
		&user.AvatarURL,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
