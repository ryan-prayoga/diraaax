package service

import (
	"context"
	"errors"

	"github.com/ryanprayoga/diraaax/backend/internal/domain"
	"github.com/ryanprayoga/diraaax/backend/internal/repository"
)

type MemoryLocationService struct {
	repository *repository.MemoryLocationRepository
}

func NewMemoryLocationService(repository *repository.MemoryLocationRepository) *MemoryLocationService {
	return &MemoryLocationService{repository: repository}
}

func (s *MemoryLocationService) List(ctx context.Context) ([]domain.MemoryLocation, error) {
	return s.repository.List(ctx)
}

func (s *MemoryLocationService) Create(
	ctx context.Context,
	input domain.CreateMemoryLocationInput,
) (*domain.MemoryLocation, error) {
	title := cleanString(input.Title)
	if title == "" {
		return nil, errors.Join(ErrInvalidInput, errors.New("title is required"))
	}

	return s.repository.Create(
		ctx,
		title,
		cleanOptionalString(input.Description),
		input.Lat,
		input.Lng,
		cleanOptionalString(input.ImageURL),
		input.RelatedMemoryID,
	)
}
