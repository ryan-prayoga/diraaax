package domain

import "time"

type MemoryLocation struct {
	ID              int64     `json:"id"`
	Title           string    `json:"title"`
	Description     *string   `json:"description,omitempty"`
	Lat             float64   `json:"lat"`
	Lng             float64   `json:"lng"`
	ImageURL        *string   `json:"image_url,omitempty"`
	RelatedMemoryID *int64    `json:"related_memory_id,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type CreateMemoryLocationInput struct {
	Title           string  `json:"title"`
	Description     *string `json:"description"`
	Lat             float64 `json:"lat"`
	Lng             float64 `json:"lng"`
	ImageURL        *string `json:"image_url"`
	RelatedMemoryID *int64  `json:"related_memory_id"`
}
