package domain

import "time"

type BucketListItem struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title"`
	Description *string    `json:"description,omitempty"`
	Category    string     `json:"category"`
	Status      string     `json:"status"`
	TargetDate  *time.Time `json:"target_date,omitempty"`
	CreatedBy   *int64     `json:"created_by,omitempty"`
	CompletedBy *int64     `json:"completed_by,omitempty"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	SortOrder   int        `json:"sort_order"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type CreateBucketListItemInput struct {
	Title       string  `json:"title"`
	Description *string `json:"description"`
	Category    string  `json:"category"`
	TargetDate  *string `json:"target_date"`
	SortOrder   *int    `json:"sort_order"`
}
