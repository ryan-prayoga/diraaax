package domain

import "time"

type Memory struct {
	ID                   int64      `json:"id"`
	Title                *string    `json:"title,omitempty"`
	Description          *string    `json:"description,omitempty"`
	ImageURL             *string    `json:"image_url,omitempty"`
	MemoryDate           *time.Time `json:"memory_date,omitempty"`
	SourceGalleryItemID  *int64     `json:"source_gallery_item_id,omitempty"`
	CreatedBy            *int64     `json:"created_by,omitempty"`
	IsFeatured           bool       `json:"is_featured"`
	IsMemoryOfDayEnabled bool       `json:"is_memory_of_day_enabled"`
	CreatedAt            time.Time  `json:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at"`
}

type CreateMemoryInput struct {
	Title                *string `json:"title"`
	Description          *string `json:"description"`
	ImageURL             *string `json:"image_url"`
	MemoryDate           *string `json:"memory_date"`
	SourceGalleryItemID  *int64  `json:"source_gallery_item_id"`
	IsFeatured           bool    `json:"is_featured"`
	IsMemoryOfDayEnabled *bool   `json:"is_memory_of_day_enabled"`
}
