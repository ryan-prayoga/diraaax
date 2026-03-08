package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ryanprayoga/diraaax/backend/internal/models"
	"github.com/ryanprayoga/diraaax/backend/internal/repositories"
	"github.com/ryanprayoga/diraaax/backend/internal/utils"
)

type AuthService struct {
	repo *repositories.AuthRepository
}

var ErrInvalidPIN = errors.New("invalid PIN")

var ErrInvalidSession = errors.New("invalid session")

func NewAuthService(repo *repositories.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

type AuthResult struct {
	Session *models.Session `json:"session"`
	User    *models.User    `json:"user,omitempty"`
}

func (s *AuthService) VerifyPIN(ctx context.Context, pin string) (*AuthResult, error) {
	ac, err := s.repo.VerifyPIN(ctx, pin)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidPIN, err)
	}

	token, err := utils.GenerateToken(32)
	if err != nil {
		return nil, fmt.Errorf("failed to generate session token: %w", err)
	}

	expiresAt := time.Now().Add(7 * 24 * time.Hour) // 7 days
	session, err := s.repo.CreateSession(ctx, token, ac.CreatedBy, &ac.ID, expiresAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	result := &AuthResult{Session: session}

	if ac.CreatedBy != nil {
		user, err := s.repo.GetUserByID(ctx, *ac.CreatedBy)
		if err == nil {
			result.User = user
		}
	}

	return result, nil
}

func (s *AuthService) GetSession(ctx context.Context, token string) (*AuthResult, error) {
	session, err := s.repo.GetSessionByToken(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidSession, err)
	}

	result := &AuthResult{Session: session}

	if session.UserID != nil {
		user, err := s.repo.GetUserByID(ctx, *session.UserID)
		if err == nil {
			result.User = user
		}
	}

	return result, nil
}

func (s *AuthService) Logout(ctx context.Context, token string) error {
	return s.repo.DeleteSession(ctx, token)
}
