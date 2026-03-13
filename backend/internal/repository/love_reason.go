package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/ryanprayoga/diraaax/backend/internal/domain"
)

type LoveReasonRepository struct {
	pool *pgxpool.Pool
}

func NewLoveReasonRepository(pool *pgxpool.Pool) *LoveReasonRepository {
	return &LoveReasonRepository{pool: pool}
}

func (r *LoveReasonRepository) List(ctx context.Context) ([]domain.LoveReason, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT id, message, created_by, visible_to, is_pinned, created_at, updated_at
		FROM love_reasons
		ORDER BY is_pinned DESC, created_at DESC, id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []domain.LoveReason
	for rows.Next() {
		var item domain.LoveReason
		if err := rows.Scan(
			&item.ID,
			&item.Message,
			&item.CreatedBy,
			&item.VisibleTo,
			&item.IsPinned,
			&item.CreatedAt,
			&item.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (r *LoveReasonRepository) Create(
	ctx context.Context,
	message string,
	createdBy *int64,
	visibleTo string,
	isPinned bool,
) (*domain.LoveReason, error) {
	var item domain.LoveReason
	err := r.pool.QueryRow(ctx, `
		INSERT INTO love_reasons (message, created_by, visible_to, is_pinned)
		VALUES ($1, $2, $3, $4)
		RETURNING id, message, created_by, visible_to, is_pinned, created_at, updated_at
	`, message, createdBy, visibleTo, isPinned).Scan(
		&item.ID,
		&item.Message,
		&item.CreatedBy,
		&item.VisibleTo,
		&item.IsPinned,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *LoveReasonRepository) Delete(ctx context.Context, id int64) error {
	commandTag, err := r.pool.Exec(ctx, `
		DELETE FROM love_reasons
		WHERE id = $1
	`, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
