package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ryanprayoga/diraaax/backend/internal/models"
)

type SecretNoteRepository struct {
	pool *pgxpool.Pool
}

func NewSecretNoteRepository(pool *pgxpool.Pool) *SecretNoteRepository {
	return &SecretNoteRepository{pool: pool}
}

func (r *SecretNoteRepository) List(ctx context.Context) ([]models.SecretNote, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, title, content, note_type, visible_to, created_by, created_at
		 FROM secret_notes
		 ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []models.SecretNote
	for rows.Next() {
		var n models.SecretNote
		if err := rows.Scan(&n.ID, &n.Title, &n.Content, &n.NoteType, &n.VisibleTo, &n.CreatedBy, &n.CreatedAt); err != nil {
			return nil, err
		}
		notes = append(notes, n)
	}
	return notes, rows.Err()
}

func (r *SecretNoteRepository) Create(ctx context.Context, title, content string, noteType, visibleTo *string, createdBy *int) (*models.SecretNote, error) {
	var n models.SecretNote
	err := r.pool.QueryRow(ctx,
		`INSERT INTO secret_notes (title, content, note_type, visible_to, created_by, created_at)
		 VALUES ($1, $2, $3, $4, $5, NOW())
		 RETURNING id, title, content, note_type, visible_to, created_by, created_at`,
		title, content, noteType, visibleTo, createdBy,
	).Scan(&n.ID, &n.Title, &n.Content, &n.NoteType, &n.VisibleTo, &n.CreatedBy, &n.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &n, nil
}
