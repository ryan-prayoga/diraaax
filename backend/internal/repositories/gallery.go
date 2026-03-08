package repositories

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ryanprayoga/diraaax/backend/internal/models"
)

type GalleryRepository struct {
	pool *pgxpool.Pool
}

func NewGalleryRepository(pool *pgxpool.Pool) *GalleryRepository {
	return &GalleryRepository{pool: pool}
}

func (r *GalleryRepository) List(ctx context.Context) ([]models.GalleryItem, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, title, caption, category, image_filename, image_path, taken_at, uploaded_by, created_at
		 FROM gallery_items
		 ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.GalleryItem
	for rows.Next() {
		var g models.GalleryItem
		if err := rows.Scan(&g.ID, &g.Title, &g.Caption, &g.Category, &g.FileName, &g.FilePath, &g.TakenAt, &g.UploadedBy, &g.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, g)
	}
	return items, rows.Err()
}

func (r *GalleryRepository) Create(ctx context.Context, title string, caption *string, category *string, fileName, filePath string, takenAt *time.Time, uploadedBy *int) (*models.GalleryItem, error) {
	var g models.GalleryItem
	err := r.pool.QueryRow(ctx,
		`INSERT INTO gallery_items (title, caption, category, image_filename, image_path, taken_at, uploaded_by)
		 VALUES ($1, $2, COALESCE($3, 'random'), $4, $5, $6, $7)
		 RETURNING id, title, caption, category, image_filename, image_path, taken_at, uploaded_by, created_at`,

		title, caption, category, fileName, filePath, takenAt, uploadedBy,
	).Scan(&g.ID, &g.Title, &g.Caption, &g.Category, &g.FileName, &g.FilePath, &g.TakenAt, &g.UploadedBy, &g.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &g, nil
}

func (r *GalleryRepository) GetByID(ctx context.Context, id int) (*models.GalleryItem, error) {
	var g models.GalleryItem
	err := r.pool.QueryRow(ctx,
		`SELECT id, title, caption, category, image_filename, image_path, taken_at, uploaded_by, created_at
		 FROM gallery_items
		 WHERE id = $1`,
		id,
	).Scan(&g.ID, &g.Title, &g.Caption, &g.Category, &g.FileName, &g.FilePath, &g.TakenAt, &g.UploadedBy, &g.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &g, nil
}

func (r *GalleryRepository) Delete(ctx context.Context, id int) error {
	_, err := r.pool.Exec(ctx, `DELETE FROM gallery_items WHERE id = $1`, id)
	return err
}
