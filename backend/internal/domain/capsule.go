package domain

import "time"

type LoveCapsule struct {
	ID            int64      `json:"id"`
	Title         string     `json:"title"`
	Message       string     `json:"message"`
	OpenDate      time.Time  `json:"open_date"`
	IsOpened      bool       `json:"is_opened"`
	OpenedAt      *time.Time `json:"opened_at,omitempty"`
	CreatedBy     *int64     `json:"created_by,omitempty"`
	VisibleTo     string     `json:"visible_to"`
	CoverImageURL *string    `json:"cover_image_url,omitempty"`
	MusicURL      *string    `json:"music_url,omitempty"`
	ThemeVariant  string     `json:"theme_variant"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

type LoveCapsuleScene struct {
	ID           int64     `json:"id"`
	CapsuleID    int64     `json:"capsule_id"`
	SceneOrder   int       `json:"scene_order"`
	SceneType    string    `json:"scene_type"`
	Title        *string   `json:"title,omitempty"`
	Content      *string   `json:"content,omitempty"`
	ImageURL     *string   `json:"image_url,omitempty"`
	AnimationKey *string   `json:"animation_key,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CreateLoveCapsuleInput struct {
	Title         string  `json:"title"`
	Message       string  `json:"message"`
	OpenDate      string  `json:"open_date"`
	VisibleTo     string  `json:"visible_to"`
	CoverImageURL *string `json:"cover_image_url"`
	MusicURL      *string `json:"music_url"`
	ThemeVariant  string  `json:"theme_variant"`
}

type CreateLoveCapsuleSceneInput struct {
	SceneOrder   int     `json:"scene_order"`
	SceneType    string  `json:"scene_type"`
	Title        *string `json:"title"`
	Content      *string `json:"content"`
	ImageURL     *string `json:"image_url"`
	AnimationKey *string `json:"animation_key"`
}
