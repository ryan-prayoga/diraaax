package domain

import "time"

type LoveReason struct {
	ID        int64     `json:"id"`
	Message   string    `json:"message"`
	CreatedBy *int64    `json:"created_by,omitempty"`
	VisibleTo string    `json:"visible_to"`
	IsPinned  bool      `json:"is_pinned"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateLoveReasonInput struct {
	Message   string `json:"message"`
	VisibleTo string `json:"visible_to"`
	IsPinned  bool   `json:"is_pinned"`
}
