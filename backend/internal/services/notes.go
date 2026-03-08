package services

import (
	"context"

	"github.com/ryanprayoga/diraaax/backend/internal/models"
	"github.com/ryanprayoga/diraaax/backend/internal/repositories"
)

type SecretNoteService struct {
	repo *repositories.SecretNoteRepository
}

func NewSecretNoteService(repo *repositories.SecretNoteRepository) *SecretNoteService {
	return &SecretNoteService{repo: repo}
}

func (s *SecretNoteService) List(ctx context.Context) ([]models.SecretNote, error) {
	notes, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	if notes == nil {
		notes = []models.SecretNote{}
	}
	return notes, nil
}

func (s *SecretNoteService) Create(ctx context.Context, title, content string, noteType, visibleTo *string, createdBy *int) (*models.SecretNote, error) {
	return s.repo.Create(ctx, title, content, noteType, visibleTo, createdBy)
}
