package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/ryanprayoga/diraaax/backend/internal/domain"
)

type VoiceNoteRepository struct {
	pool *pgxpool.Pool
}

func NewVoiceNoteRepository(pool *pgxpool.Pool) *VoiceNoteRepository {
	return &VoiceNoteRepository{pool: pool}
}

func (r *VoiceNoteRepository) List(ctx context.Context) ([]domain.VoiceNote, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT id, title, audio_url, duration_seconds, transcript, created_by, created_at, updated_at
		FROM voice_notes
		ORDER BY created_at DESC, id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []domain.VoiceNote
	for rows.Next() {
		var item domain.VoiceNote
		if err := rows.Scan(
			&item.ID,
			&item.Title,
			&item.AudioURL,
			&item.DurationSeconds,
			&item.Transcript,
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

func (r *VoiceNoteRepository) Create(
	ctx context.Context,
	title *string,
	audioURL string,
	durationSeconds *int,
	transcript *string,
	createdBy *int64,
) (*domain.VoiceNote, error) {
	var item domain.VoiceNote
	err := r.pool.QueryRow(ctx, `
		INSERT INTO voice_notes (title, audio_url, duration_seconds, transcript, created_by)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, title, audio_url, duration_seconds, transcript, created_by, created_at, updated_at
	`, title, audioURL, durationSeconds, transcript, createdBy).Scan(
		&item.ID,
		&item.Title,
		&item.AudioURL,
		&item.DurationSeconds,
		&item.Transcript,
		&item.CreatedBy,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &item, nil
}
