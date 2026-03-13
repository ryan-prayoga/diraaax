package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/ryanprayoga/diraaax/backend/internal/domain"
)

type BucketListRepository struct {
	pool *pgxpool.Pool
}

func NewBucketListRepository(pool *pgxpool.Pool) *BucketListRepository {
	return &BucketListRepository{pool: pool}
}

func (r *BucketListRepository) List(ctx context.Context) ([]domain.BucketListItem, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT id, title, description, category, status, target_date, created_by,
		       completed_by, completed_at, sort_order, created_at, updated_at
		FROM bucket_list_items
		ORDER BY sort_order ASC, created_at DESC, id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []domain.BucketListItem
	for rows.Next() {
		var item domain.BucketListItem
		if err := rows.Scan(
			&item.ID,
			&item.Title,
			&item.Description,
			&item.Category,
			&item.Status,
			&item.TargetDate,
			&item.CreatedBy,
			&item.CompletedBy,
			&item.CompletedAt,
			&item.SortOrder,
			&item.CreatedAt,
			&item.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (r *BucketListRepository) Create(
	ctx context.Context,
	title string,
	description *string,
	category string,
	targetDate any,
	createdBy *int64,
	sortOrder int,
) (*domain.BucketListItem, error) {
	var item domain.BucketListItem
	err := r.pool.QueryRow(ctx, `
		INSERT INTO bucket_list_items (title, description, category, target_date, created_by, sort_order)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, title, description, category, status, target_date, created_by,
		          completed_by, completed_at, sort_order, created_at, updated_at
	`, title, description, category, targetDate, createdBy, sortOrder).Scan(
		&item.ID,
		&item.Title,
		&item.Description,
		&item.Category,
		&item.Status,
		&item.TargetDate,
		&item.CreatedBy,
		&item.CompletedBy,
		&item.CompletedAt,
		&item.SortOrder,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *BucketListRepository) Toggle(ctx context.Context, id int64, completedBy *int64) (*domain.BucketListItem, error) {
	var item domain.BucketListItem
	err := r.pool.QueryRow(ctx, `
		UPDATE bucket_list_items
		SET status = CASE WHEN status = 'done' THEN 'pending' ELSE 'done' END,
		    completed_by = CASE WHEN status = 'done' THEN NULL ELSE $2 END,
		    completed_at = CASE WHEN status = 'done' THEN NULL ELSE NOW() END,
		    updated_at = NOW()
		WHERE id = $1
		RETURNING id, title, description, category, status, target_date, created_by,
		          completed_by, completed_at, sort_order, created_at, updated_at
	`, id, completedBy).Scan(
		&item.ID,
		&item.Title,
		&item.Description,
		&item.Category,
		&item.Status,
		&item.TargetDate,
		&item.CreatedBy,
		&item.CompletedBy,
		&item.CompletedAt,
		&item.SortOrder,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *BucketListRepository) Delete(ctx context.Context, id int64) error {
	commandTag, err := r.pool.Exec(ctx, `
		DELETE FROM bucket_list_items
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
