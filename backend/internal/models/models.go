package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Nickname  string    `json:"nickname"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type AccessCode struct {
	ID        int    `json:"id"`
	Label     string `json:"label"`
	CreatedBy *int   `json:"created_by"`
}

type Session struct {
	ID           int       `json:"id"`
	Token        string    `json:"token"`
	UserID       *int      `json:"user_id"`
	AccessCodeID *int      `json:"access_code_id,omitempty"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type GalleryItem struct {
	ID         int        `json:"id"`
	Title      string     `json:"title"`
	Caption    *string    `json:"caption"`
	Category   *string    `json:"category"`
	FileName   string     `json:"file_name"`
	FilePath   string     `json:"file_path"`
	TakenAt    *time.Time `json:"taken_at"`
	UploadedBy *int       `json:"uploaded_by"`
	CreatedAt  time.Time  `json:"created_at"`
}

type FuturePlan struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Category    *string    `json:"category"`
	Status      string     `json:"status"`
	CreatedBy   *int       `json:"created_by"`
	CompletedAt *time.Time `json:"completed_at"`
	CreatedAt   time.Time  `json:"created_at"`
}

type SecretNote struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	NoteType  *string   `json:"note_type"`
	VisibleTo *string   `json:"visible_to"`
	CreatedBy *int      `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}
