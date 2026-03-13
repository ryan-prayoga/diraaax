package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/ryanprayoga/diraaax/backend/internal/domain"
)

type TimelineRepository struct {
	pool *pgxpool.Pool
}

func NewTimelineRepository(pool *pgxpool.Pool) *TimelineRepository {
	return &TimelineRepository{pool: pool}
}

func (r *TimelineRepository) List(ctx context.Context) ([]domain.TimelineEvent, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT id, title, description, event_date, image_url, event_type, created_by, created_at, updated_at
		FROM timeline_events
		ORDER BY event_date ASC, id ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []domain.TimelineEvent
	for rows.Next() {
		var item domain.TimelineEvent
		if err := rows.Scan(
			&item.ID,
			&item.Title,
			&item.Description,
			&item.EventDate,
			&item.ImageURL,
			&item.EventType,
			&item.CreatedBy,
			&item.CreatedAt,
			&item.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

func (r *TimelineRepository) Create(
	ctx context.Context,
	title string,
	description *string,
	eventDate any,
	imageURL *string,
	eventType string,
	createdBy *int64,
) (*domain.TimelineEvent, error) {
	var item domain.TimelineEvent
	err := r.pool.QueryRow(ctx, `
		INSERT INTO timeline_events (title, description, event_date, image_url, event_type, created_by)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, title, description, event_date, image_url, event_type, created_by, created_at, updated_at
	`, title, description, eventDate, imageURL, eventType, createdBy).Scan(
		&item.ID,
		&item.Title,
		&item.Description,
		&item.EventDate,
		&item.ImageURL,
		&item.EventType,
		&item.CreatedBy,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *TimelineRepository) Delete(ctx context.Context, id int64) error {
	commandTag, err := r.pool.Exec(ctx, `
		DELETE FROM timeline_events
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
