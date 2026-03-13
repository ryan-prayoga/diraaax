package domain

import "time"

type TimelineEvent struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description *string   `json:"description,omitempty"`
	EventDate   time.Time `json:"event_date"`
	ImageURL    *string   `json:"image_url,omitempty"`
	EventType   string    `json:"event_type"`
	CreatedBy   *int64    `json:"created_by,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateTimelineEventInput struct {
	Title       string  `json:"title"`
	Description *string `json:"description"`
	EventDate   string  `json:"event_date"`
	ImageURL    *string `json:"image_url"`
	EventType   string  `json:"event_type"`
}
