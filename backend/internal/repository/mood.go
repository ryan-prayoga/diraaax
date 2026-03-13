package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/ryanprayoga/diraaax/backend/internal/domain"
)

type MoodRepository struct {
	pool *pgxpool.Pool
}

func NewMoodRepository(pool *pgxpool.Pool) *MoodRepository {
	return &MoodRepository{pool: pool}
}

func (r *MoodRepository) List(ctx context.Context) ([]domain.DailyMood, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT
			m.id,
			m.user_id,
			m.mood,
			m.note,
			m.mood_date,
			m.created_at,
			u.id,
			u.display_name,
			u.nickname,
			u.slug,
			u.avatar_url,
			u.is_active,
			u.created_at,
			u.updated_at
		FROM daily_moods m
		JOIN users u ON u.id = m.user_id
		ORDER BY m.mood_date DESC, m.created_at DESC, m.id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []domain.DailyMood
	for rows.Next() {
		var item domain.DailyMood
		var user domain.User
		if err := rows.Scan(
			&item.ID,
			&item.UserID,
			&item.Mood,
			&item.Note,
			&item.MoodDate,
			&item.CreatedAt,
			&user.ID,
			&user.DisplayName,
			&user.Nickname,
			&user.Slug,
			&user.AvatarURL,
			&user.IsActive,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return nil, err
		}
		item.User = &user
		items = append(items, item)
	}
	return items, rows.Err()
}

func (r *MoodRepository) Create(
	ctx context.Context,
	userID int64,
	mood string,
	note *string,
	moodDate any,
) (*domain.DailyMood, error) {
	var item domain.DailyMood
	var user domain.User
	err := r.pool.QueryRow(ctx, `
		WITH inserted AS (
			INSERT INTO daily_moods (user_id, mood, note, mood_date)
			VALUES ($1, $2, $3, $4)
			RETURNING id, user_id, mood, note, mood_date, created_at
		)
		SELECT
			i.id,
			i.user_id,
			i.mood,
			i.note,
			i.mood_date,
			i.created_at,
			u.id,
			u.display_name,
			u.nickname,
			u.slug,
			u.avatar_url,
			u.is_active,
			u.created_at,
			u.updated_at
		FROM inserted i
		JOIN users u ON u.id = i.user_id
	`, userID, mood, note, moodDate).Scan(
		&item.ID,
		&item.UserID,
		&item.Mood,
		&item.Note,
		&item.MoodDate,
		&item.CreatedAt,
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
	item.User = &user
	return &item, nil
}
