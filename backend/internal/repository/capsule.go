package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/ryanprayoga/diraaax/backend/internal/domain"
)

type CapsuleRepository struct {
	pool *pgxpool.Pool
}

func NewCapsuleRepository(pool *pgxpool.Pool) *CapsuleRepository {
	return &CapsuleRepository{pool: pool}
}

func (r *CapsuleRepository) List(ctx context.Context) ([]domain.LoveCapsule, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT id, title, message, open_date, is_opened, opened_at, created_by,
		       visible_to, cover_image_url, music_url, theme_variant, created_at, updated_at
		FROM love_capsules
		ORDER BY open_date DESC, id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []domain.LoveCapsule
	for rows.Next() {
		var item domain.LoveCapsule
		if err := rows.Scan(
			&item.ID,
			&item.Title,
			&item.Message,
			&item.OpenDate,
			&item.IsOpened,
			&item.OpenedAt,
			&item.CreatedBy,
			&item.VisibleTo,
			&item.CoverImageURL,
			&item.MusicURL,
			&item.ThemeVariant,
			&item.CreatedAt,
			&item.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (r *CapsuleRepository) GetByID(ctx context.Context, id int64) (*domain.LoveCapsule, error) {
	var item domain.LoveCapsule
	err := r.pool.QueryRow(ctx, `
		SELECT id, title, message, open_date, is_opened, opened_at, created_by,
		       visible_to, cover_image_url, music_url, theme_variant, created_at, updated_at
		FROM love_capsules
		WHERE id = $1
	`, id).Scan(
		&item.ID,
		&item.Title,
		&item.Message,
		&item.OpenDate,
		&item.IsOpened,
		&item.OpenedAt,
		&item.CreatedBy,
		&item.VisibleTo,
		&item.CoverImageURL,
		&item.MusicURL,
		&item.ThemeVariant,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *CapsuleRepository) Create(
	ctx context.Context,
	title string,
	message string,
	openDate any,
	createdBy *int64,
	visibleTo string,
	coverImageURL *string,
	musicURL *string,
	themeVariant string,
) (*domain.LoveCapsule, error) {
	var item domain.LoveCapsule
	err := r.pool.QueryRow(ctx, `
		INSERT INTO love_capsules (
			title,
			message,
			open_date,
			created_by,
			visible_to,
			cover_image_url,
			music_url,
			theme_variant
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, title, message, open_date, is_opened, opened_at, created_by,
		          visible_to, cover_image_url, music_url, theme_variant, created_at, updated_at
	`, title, message, openDate, createdBy, visibleTo, coverImageURL, musicURL, themeVariant).Scan(
		&item.ID,
		&item.Title,
		&item.Message,
		&item.OpenDate,
		&item.IsOpened,
		&item.OpenedAt,
		&item.CreatedBy,
		&item.VisibleTo,
		&item.CoverImageURL,
		&item.MusicURL,
		&item.ThemeVariant,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *CapsuleRepository) Open(ctx context.Context, id int64) (*domain.LoveCapsule, error) {
	var item domain.LoveCapsule
	err := r.pool.QueryRow(ctx, `
		UPDATE love_capsules
		SET is_opened = true,
		    opened_at = COALESCE(opened_at, NOW()),
		    updated_at = NOW()
		WHERE id = $1
		RETURNING id, title, message, open_date, is_opened, opened_at, created_by,
		          visible_to, cover_image_url, music_url, theme_variant, created_at, updated_at
	`, id).Scan(
		&item.ID,
		&item.Title,
		&item.Message,
		&item.OpenDate,
		&item.IsOpened,
		&item.OpenedAt,
		&item.CreatedBy,
		&item.VisibleTo,
		&item.CoverImageURL,
		&item.MusicURL,
		&item.ThemeVariant,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *CapsuleRepository) ListScenes(ctx context.Context, capsuleID int64) ([]domain.LoveCapsuleScene, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT id, capsule_id, scene_order, scene_type, title, content, image_url, animation_key, created_at, updated_at
		FROM love_capsule_scenes
		WHERE capsule_id = $1
		ORDER BY scene_order ASC, id ASC
	`, capsuleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []domain.LoveCapsuleScene
	for rows.Next() {
		var item domain.LoveCapsuleScene
		if err := rows.Scan(
			&item.ID,
			&item.CapsuleID,
			&item.SceneOrder,
			&item.SceneType,
			&item.Title,
			&item.Content,
			&item.ImageURL,
			&item.AnimationKey,
			&item.CreatedAt,
			&item.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (r *CapsuleRepository) CreateScene(
	ctx context.Context,
	capsuleID int64,
	sceneOrder int,
	sceneType string,
	title *string,
	content *string,
	imageURL *string,
	animationKey *string,
) (*domain.LoveCapsuleScene, error) {
	var item domain.LoveCapsuleScene
	err := r.pool.QueryRow(ctx, `
		INSERT INTO love_capsule_scenes (
			capsule_id,
			scene_order,
			scene_type,
			title,
			content,
			image_url,
			animation_key
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, capsule_id, scene_order, scene_type, title, content, image_url, animation_key, created_at, updated_at
	`, capsuleID, sceneOrder, sceneType, title, content, imageURL, animationKey).Scan(
		&item.ID,
		&item.CapsuleID,
		&item.SceneOrder,
		&item.SceneType,
		&item.Title,
		&item.Content,
		&item.ImageURL,
		&item.AnimationKey,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &item, nil
}
