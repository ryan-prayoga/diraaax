package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/ryanprayoga/diraaax/backend/internal/domain"
	"github.com/ryanprayoga/diraaax/backend/internal/repository"
)

type BucketListService struct {
	repository *repository.BucketListRepository
}

func NewBucketListService(repository *repository.BucketListRepository) *BucketListService {
	return &BucketListService{repository: repository}
}

func (s *BucketListService) List(ctx context.Context) ([]domain.BucketListItem, error) {
	return s.repository.List(ctx)
}

func (s *BucketListService) Create(
	ctx context.Context,
	createdBy *int64,
	input domain.CreateBucketListItemInput,
) (*domain.BucketListItem, error) {
	title := cleanString(input.Title)
	if title == "" {
		return nil, errors.Join(ErrInvalidInput, errors.New("title is required"))
	}

	targetDate, err := parseOptionalDate(input.TargetDate)
	if err != nil {
		return nil, errors.Join(ErrInvalidInput, fmt.Errorf("target_date must use YYYY-MM-DD: %w", err))
	}

	category := cleanString(input.Category)
	if category == "" {
		category = "random"
	}

	sortOrder := 0
	if input.SortOrder != nil {
		sortOrder = *input.SortOrder
	}

	return s.repository.Create(
		ctx,
		title,
		cleanOptionalString(input.Description),
		category,
		targetDate,
		createdBy,
		sortOrder,
	)
}

func (s *BucketListService) Toggle(ctx context.Context, id int64, completedBy *int64) (*domain.BucketListItem, error) {
	if id <= 0 {
		return nil, errors.Join(ErrInvalidInput, errors.New("invalid bucket list item id"))
	}
	item, err := s.repository.Toggle(ctx, id, completedBy)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.Join(ErrNotFound, errors.New("bucket list item not found"))
		}
		return nil, err
	}
	return item, nil
}

func (s *BucketListService) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.Join(ErrInvalidInput, errors.New("invalid bucket list item id"))
	}
	if err := s.repository.Delete(ctx, id); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errors.Join(ErrNotFound, errors.New("bucket list item not found"))
		}
		return err
	}
	return nil
}
