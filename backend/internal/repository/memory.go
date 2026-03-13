package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/ryanprayoga/diraaax/backend/internal/domain"
)

type MemoryRepository struct {
	pool *pgxpool.Pool
}

func NewMemoryRepository(pool *pgxpool.Pool) *MemoryRepository {
	return &MemoryRepository{pool: pool}
}

func (r *MemoryRepository) List(ctx context.Context) ([]domain.Memory, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT id, title, description, image_url, memory_date, source_gallery_item_id, created_by,
		       is_featured, is_memory_of_day_enabled, created_at, updated_at
		FROM memories
		ORDER BY COALESCE(memory_date, created_at::date) DESC, id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []domain.Memory
	for rows.Next() {
		var item domain.Memory
		if err := rows.Scan(
			&item.ID,
			&item.Title,
			&item.Description,
			&item.ImageURL,
			&item.MemoryDate,
			&item.SourceGalleryItemID,
			&item.CreatedBy,
			&item.IsFeatured,
			&item.IsMemoryOfDayEnabled,
			&item.CreatedAt,
			&item.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (r *MemoryRepository) Random(ctx context.Context) (*domain.Memory, error) {
	var item domain.Memory
	err := r.pool.QueryRow(ctx, `
		SELECT id, title, description, image_url, memory_date, source_gallery_item_id, created_by,
		       is_featured, is_memory_of_day_enabled, created_at, updated_at
		FROM memories
		WHERE is_memory_of_day_enabled = true
		ORDER BY random()
		LIMIT 1
	`).Scan(
		&item.ID,
		&item.Title,
		&item.Description,
		&item.ImageURL,
		&item.MemoryDate,
		&item.SourceGalleryItemID,
		&item.CreatedBy,
		&item.IsFeatured,
		&item.IsMemoryOfDayEnabled,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MemoryRepository) Create(
	ctx context.Context,
	title *string,
	description *string,
	imageURL *string,
	memoryDate any,
	sourceGalleryItemID *int64,
	createdBy *int64,
	isFeatured bool,
	isMemoryOfDayEnabled bool,
) (*domain.Memory, error) {
	var item domain.Memory
	err := r.pool.QueryRow(ctx, `
		INSERT INTO memories (
			title,
			description,
			image_url,
			memory_date,
			source_gallery_item_id,
			created_by,
			is_featured,
			is_memory_of_day_enabled
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, title, description, image_url, memory_date, source_gallery_item_id, created_by,
		          is_featured, is_memory_of_day_enabled, created_at, updated_at
	`, title, description, imageURL, memoryDate, sourceGalleryItemID, createdBy, isFeatured, isMemoryOfDayEnabled).Scan(
		&item.ID,
		&item.Title,
		&item.Description,
		&item.ImageURL,
		&item.MemoryDate,
		&item.SourceGalleryItemID,
		&item.CreatedBy,
		&item.IsFeatured,
		&item.IsMemoryOfDayEnabled,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MemoryRepository) Delete(ctx context.Context, id int64) error {
	commandTag, err := r.pool.Exec(ctx, `
		DELETE FROM memories
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
