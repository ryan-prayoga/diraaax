package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/ryanprayoga/diraaax/backend/internal/domain"
)

type MemoryLocationRepository struct {
	pool *pgxpool.Pool
}

func NewMemoryLocationRepository(pool *pgxpool.Pool) *MemoryLocationRepository {
	return &MemoryLocationRepository{pool: pool}
}

func (r *MemoryLocationRepository) List(ctx context.Context) ([]domain.MemoryLocation, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT id, title, description, lat, lng, image_url, related_memory_id, created_at, updated_at
		FROM memory_locations
		ORDER BY created_at DESC, id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []domain.MemoryLocation
	for rows.Next() {
		var item domain.MemoryLocation
		if err := rows.Scan(
			&item.ID,
			&item.Title,
			&item.Description,
			&item.Lat,
			&item.Lng,
			&item.ImageURL,
			&item.RelatedMemoryID,
			&item.CreatedAt,
			&item.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (r *MemoryLocationRepository) Create(
	ctx context.Context,
	title string,
	description *string,
	lat float64,
	lng float64,
	imageURL *string,
	relatedMemoryID *int64,
) (*domain.MemoryLocation, error) {
	var item domain.MemoryLocation
	err := r.pool.QueryRow(ctx, `
		INSERT INTO memory_locations (title, description, lat, lng, image_url, related_memory_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, title, description, lat, lng, image_url, related_memory_id, created_at, updated_at
	`, title, description, lat, lng, imageURL, relatedMemoryID).Scan(
		&item.ID,
		&item.Title,
		&item.Description,
		&item.Lat,
		&item.Lng,
		&item.ImageURL,
		&item.RelatedMemoryID,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &item, nil
}
