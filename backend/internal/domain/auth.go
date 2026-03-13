package domain

import "time"

type User struct {
	ID          int64     `json:"id"`
	DisplayName string    `json:"display_name"`
	Nickname    *string   `json:"nickname,omitempty"`
	Slug        *string   `json:"slug,omitempty"`
	AvatarURL   *string   `json:"avatar_url,omitempty"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type AccessCode struct {
	ID        int64   `json:"id"`
	Label     string  `json:"label"`
	CreatedBy *int64  `json:"created_by,omitempty"`
	CodeType  string  `json:"code_type"`
	CodeHint  *string `json:"code_hint,omitempty"`
}

type Session struct {
	ID           int64      `json:"id"`
	UserID       *int64     `json:"user_id,omitempty"`
	AccessCodeID *int64     `json:"access_code_id,omitempty"`
	ExpiresAt    time.Time  `json:"expires_at"`
	LastSeenAt   *time.Time `json:"last_seen_at,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
}

type AuthSession struct {
	Session    *Session    `json:"session"`
	User       *User       `json:"user,omitempty"`
	AccessCode *AccessCode `json:"access_code,omitempty"`
}

type VerifyPINInput struct {
	PIN string `json:"pin"`
}

func (a *AuthSession) ActorUserID() *int64 {
	if a == nil || a.Session == nil {
		return nil
	}
	if a.Session.UserID != nil {
		return a.Session.UserID
	}
	if a.AccessCode != nil {
		return a.AccessCode.CreatedBy
	}
	return nil
}
