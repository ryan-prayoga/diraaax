package domain

import "time"

type DailyMood struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Mood      string    `json:"mood"`
	Note      *string   `json:"note,omitempty"`
	MoodDate  time.Time `json:"mood_date"`
	CreatedAt time.Time `json:"created_at"`
	User      *User     `json:"user,omitempty"`
}

type CreateDailyMoodInput struct {
	Mood     string  `json:"mood"`
	Note     *string `json:"note"`
	MoodDate *string `json:"mood_date"`
}
