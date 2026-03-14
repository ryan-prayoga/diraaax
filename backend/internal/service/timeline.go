package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/ryanprayoga/diraaax/backend/internal/domain"
	"github.com/ryanprayoga/diraaax/backend/internal/repository"
)

type TimelineService struct {
	repository *repository.TimelineRepository
}

func NewTimelineService(repository *repository.TimelineRepository) *TimelineService {
	return &TimelineService{repository: repository}
}

func (s *TimelineService) List(ctx context.Context) ([]domain.TimelineEvent, error) {
	return s.repository.List(ctx)
}

func (s *TimelineService) Create(
	ctx context.Context,
	createdBy *int64,
	input domain.CreateTimelineEventInput,
) (*domain.TimelineEvent, error) {
	title := cleanString(input.Title)
	if title == "" {
		return nil, errors.Join(ErrInvalidInput, errors.New("title is required"))
	}

	eventDate, err := parseRequiredDate(input.EventDate)
	if err != nil {
		return nil, errors.Join(ErrInvalidInput, fmt.Errorf("event_date must use YYYY-MM-DD: %w", err))
	}

	eventType := cleanString(input.EventType)
	if eventType == "" {
		return nil, errors.Join(ErrInvalidInput, errors.New("event_type is required"))
	}

	return s.repository.Create(
		ctx,
		title,
		cleanOptionalString(input.Description),
		eventDate,
		cleanOptionalString(input.ImageURL),
		eventType,
		createdBy,
	)
}

func (s *TimelineService) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.Join(ErrInvalidInput, errors.New("invalid timeline id"))
	}
	if err := s.repository.Delete(ctx, id); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errors.Join(ErrNotFound, errors.New("timeline event not found"))
		}
		return err
	}
	return nil
}
