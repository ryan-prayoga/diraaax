package domain

import "time"

type VoiceNote struct {
	ID              int64     `json:"id"`
	Title           *string   `json:"title,omitempty"`
	AudioURL        string    `json:"audio_url"`
	DurationSeconds *int      `json:"duration_seconds,omitempty"`
	Transcript      *string   `json:"transcript,omitempty"`
	CreatedBy       *int64    `json:"created_by,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type CreateVoiceNoteInput struct {
	Title           *string `json:"title"`
	AudioURL        string  `json:"audio_url"`
	DurationSeconds *int    `json:"duration_seconds"`
	Transcript      *string `json:"transcript"`
}
