package repositories

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ryanprayoga/diraaax/backend/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepository struct {
	pool *pgxpool.Pool
}

func NewAuthRepository(pool *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{pool: pool}
}

func (r *AuthRepository) VerifyPIN(ctx context.Context, pin string) (*models.AccessCode, error) {
	pin = strings.TrimSpace(pin)
	if pin == "" {
		return nil, errors.New("pin is required")
	}

	rows, err := r.pool.Query(ctx,
		`SELECT id, label, created_by, code_hash
		 FROM access_codes
		 WHERE is_active = true
		 ORDER BY id DESC`,
	)
	if err != nil {
		log.Printf("[AUTH] PIN verification failed: unable to query access codes (db error: %v)", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var ac models.AccessCode
		var codeHash string

		if err := rows.Scan(&ac.ID, &ac.Label, &ac.CreatedBy, &codeHash); err != nil {
			log.Printf("[AUTH] PIN verification failed: unable to scan access code row (db error: %v)", err)
			return nil, err
		}

		matches, err := r.matchesPIN(ctx, pin, codeHash)
		if err != nil {
			log.Printf("[AUTH] Skipping invalid access code hash for access_code_id=%d: %v", ac.ID, err)
			continue
		}

		if matches {
			log.Printf("[AUTH] PIN verified successfully, access_code_id=%d", ac.ID)
			return &ac, nil
		}
	}

	if err := rows.Err(); err != nil {
		log.Printf("[AUTH] PIN verification failed: row iteration error (db error: %v)", err)
		return nil, err
	}

	log.Printf("[AUTH] PIN verification failed: no matching access code")
	return nil, errors.New("invalid pin")
}

func (r *AuthRepository) matchesPIN(ctx context.Context, pin, codeHash string) (bool, error) {
	switch {
	case strings.HasPrefix(codeHash, "$2a$"), strings.HasPrefix(codeHash, "$2b$"), strings.HasPrefix(codeHash, "$2y$"):
		err := bcrypt.CompareHashAndPassword([]byte(codeHash), []byte(pin))
		if err == nil {
			return true, nil
		}
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil
		}
		return false, err
	default:
		var matches bool
		err := r.pool.QueryRow(ctx, `SELECT crypt($1, $2) = $2`, pin, codeHash).Scan(&matches)
		if err != nil {
			return false, err
		}
		return matches, nil
	}
}

func (r *AuthRepository) CreateSession(ctx context.Context, token string, userID *int, accessCodeID *int, expiresAt time.Time) (*models.Session, error) {
	var s models.Session
	tokenHash := hashToken(token)
	err := r.pool.QueryRow(ctx,
		`INSERT INTO sessions (token_hash, user_id, access_code_id, expires_at, created_at)
		 VALUES ($1, $2, $3, $4, NOW())
		 RETURNING id, user_id, access_code_id, expires_at, created_at`,
		tokenHash, userID, accessCodeID, expiresAt,
	).Scan(&s.ID, &s.UserID, &s.AccessCodeID, &s.ExpiresAt, &s.CreatedAt)
	if err != nil {
		return nil, err
	}
	s.Token = token
	return &s, nil
}

func (r *AuthRepository) GetSessionByToken(ctx context.Context, token string) (*models.Session, error) {
	var s models.Session
	tokenHash := hashToken(token)
	err := r.pool.QueryRow(ctx,
		`SELECT id, user_id, access_code_id, expires_at, created_at
		 FROM sessions
		 WHERE token_hash = $1 AND expires_at > NOW()
		 LIMIT 1`,
		tokenHash,
	).Scan(&s.ID, &s.UserID, &s.AccessCodeID, &s.ExpiresAt, &s.CreatedAt)
	if err != nil {
		return nil, err
	}
	s.Token = token
	return &s, nil
}

func (r *AuthRepository) DeleteSession(ctx context.Context, token string) error {
	_, err := r.pool.Exec(ctx, `DELETE FROM sessions WHERE token_hash = $1`, hashToken(token))
	return err
}

func (r *AuthRepository) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	var u models.User
	err := r.pool.QueryRow(ctx,
		`SELECT id, display_name, nickname, '' AS role, created_at
		 FROM users
		 WHERE id = $1 AND is_active = true`,
		id,
	).Scan(&u.ID, &u.Name, &u.Nickname, &u.Role, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func hashToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}
