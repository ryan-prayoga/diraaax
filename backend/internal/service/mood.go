package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/ryanprayoga/diraaax/backend/internal/domain"
	"github.com/ryanprayoga/diraaax/backend/internal/repository"
)

type MoodService struct {
	repository *repository.MoodRepository
}

func NewMoodService(repository *repository.MoodRepository) *MoodService {
	return &MoodService{repository: repository}
}

func (s *MoodService) List(ctx context.Context) ([]domain.DailyMood, error) {
	return s.repository.List(ctx)
}

func (s *MoodService) Create(
	ctx context.Context,
	userID *int64,
	input domain.CreateDailyMoodInput,
) (*domain.DailyMood, error) {
	if userID == nil {
		return nil, errors.Join(ErrForbidden, errors.New("mood entries require a user-bound session"))
	}

	mood := cleanString(input.Mood)
	if mood == "" {
		return nil, errors.Join(ErrInvalidInput, errors.New("mood is required"))
	}

	var moodDate any = nil
	if input.MoodDate != nil {
		parsed, err := parseOptionalDate(input.MoodDate)
		if err != nil {
			return nil, errors.Join(ErrInvalidInput, fmt.Errorf("mood_date must use YYYY-MM-DD: %w", err))
		}
		moodDate = parsed
	}

	return s.repository.Create(
		ctx,
		*userID,
		mood,
		cleanOptionalString(input.Note),
		moodDate,
	)
}
