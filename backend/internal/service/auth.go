package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/ryanprayoga/diraaax/backend/internal/domain"
	"github.com/ryanprayoga/diraaax/backend/internal/repository"
	"github.com/ryanprayoga/diraaax/backend/internal/utils"
)

type AuthService struct {
	repository *repository.AuthRepository
	sessionTTL time.Duration
}

func NewAuthService(repository *repository.AuthRepository, sessionTTL time.Duration) *AuthService {
	return &AuthService{
		repository: repository,
		sessionTTL: sessionTTL,
	}
}

func (s *AuthService) VerifyPIN(
	ctx context.Context,
	pin string,
	userAgent *string,
	ipAddress *string,
) (*domain.AuthSession, string, error) {
	if strings.TrimSpace(pin) == "" {
		return nil, "", errors.Join(ErrInvalidInput, errors.New("pin is required"))
	}

	token, err := utils.GenerateSecureToken(32)
	if err != nil {
		return nil, "", err
	}

	authSession, err := s.repository.VerifyPINAndCreateSession(
		ctx,
		pin,
		token,
		userAgent,
		ipAddress,
		time.Now().Add(s.sessionTTL),
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, "", errors.Join(ErrUnauthorized, errors.New("invalid pin"))
		}
		return nil, "", err
	}

	return authSession, token, nil
}

func (s *AuthService) GetSession(ctx context.Context, token string) (*domain.AuthSession, error) {
	if strings.TrimSpace(token) == "" {
		return nil, errors.Join(ErrUnauthorized, errors.New("missing session token"))
	}

	authSession, err := s.repository.GetSession(ctx, token)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.Join(ErrUnauthorized, errors.New("session not found"))
		}
		return nil, err
	}
	return authSession, nil
}

func (s *AuthService) Logout(ctx context.Context, token string) error {
	if strings.TrimSpace(token) == "" {
		return errors.Join(ErrUnauthorized, errors.New("missing session token"))
	}

	if err := s.repository.DeleteSession(ctx, token); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errors.Join(ErrUnauthorized, errors.New("session not found"))
		}
		return err
	}
	return nil
}
