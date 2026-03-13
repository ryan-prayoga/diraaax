package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/ryanprayoga/diraaax/backend/internal/domain"
	"github.com/ryanprayoga/diraaax/backend/internal/repository"
)

type MemoryService struct {
	repository *repository.MemoryRepository
}

func NewMemoryService(repository *repository.MemoryRepository) *MemoryService {
	return &MemoryService{repository: repository}
}

func (s *MemoryService) List(ctx context.Context) ([]domain.Memory, error) {
	return s.repository.List(ctx)
}

func (s *MemoryService) Random(ctx context.Context) (*domain.Memory, error) {
	item, err := s.repository.Random(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.Join(ErrNotFound, errors.New("no memories available"))
		}
		return nil, err
	}
	return item, nil
}

func (s *MemoryService) Create(
	ctx context.Context,
	createdBy *int64,
	input domain.CreateMemoryInput,
) (*domain.Memory, error) {
	title := cleanOptionalString(input.Title)
	description := cleanOptionalString(input.Description)
	imageURL := cleanOptionalString(input.ImageURL)

	if title == nil && description == nil && imageURL == nil {
		return nil, errors.Join(ErrInvalidInput, errors.New("at least one of title, description, or image_url is required"))
	}

	memoryDate, err := parseOptionalDate(input.MemoryDate)
	if err != nil {
		return nil, errors.Join(ErrInvalidInput, fmt.Errorf("memory_date must use YYYY-MM-DD: %w", err))
	}

	isMemoryOfDayEnabled := true
	if input.IsMemoryOfDayEnabled != nil {
		isMemoryOfDayEnabled = *input.IsMemoryOfDayEnabled
	}

	return s.repository.Create(
		ctx,
		title,
		description,
		imageURL,
		memoryDate,
		input.SourceGalleryItemID,
		createdBy,
		input.IsFeatured,
		isMemoryOfDayEnabled,
	)
}

func (s *MemoryService) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.Join(ErrInvalidInput, errors.New("invalid memory id"))
	}
	if err := s.repository.Delete(ctx, id); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errors.Join(ErrNotFound, errors.New("memory not found"))
		}
		return err
	}
	return nil
}
