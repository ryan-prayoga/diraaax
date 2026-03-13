package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/ryanprayoga/diraaax/backend/internal/domain"
	"github.com/ryanprayoga/diraaax/backend/internal/repository"
)

type CapsuleService struct {
	repository *repository.CapsuleRepository
}

func NewCapsuleService(repository *repository.CapsuleRepository) *CapsuleService {
	return &CapsuleService{repository: repository}
}

func (s *CapsuleService) List(ctx context.Context) ([]domain.LoveCapsule, error) {
	return s.repository.List(ctx)
}

func (s *CapsuleService) Get(ctx context.Context, id int64) (*domain.LoveCapsule, error) {
	if id <= 0 {
		return nil, errors.Join(ErrInvalidInput, errors.New("invalid capsule id"))
	}
	item, err := s.repository.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.Join(ErrNotFound, errors.New("capsule not found"))
		}
		return nil, err
	}
	return item, nil
}

func (s *CapsuleService) Create(
	ctx context.Context,
	createdBy *int64,
	input domain.CreateLoveCapsuleInput,
) (*domain.LoveCapsule, error) {
	title := cleanString(input.Title)
	if title == "" {
		return nil, errors.Join(ErrInvalidInput, errors.New("title is required"))
	}

	message := cleanString(input.Message)
	if message == "" {
		return nil, errors.Join(ErrInvalidInput, errors.New("message is required"))
	}

	openDate, err := parseDateTime(input.OpenDate)
	if err != nil {
		return nil, errors.Join(ErrInvalidInput, fmt.Errorf("open_date must use RFC3339 or YYYY-MM-DD: %w", err))
	}

	visibleTo := cleanString(input.VisibleTo)
	if visibleTo == "" {
		visibleTo = "both"
	}

	themeVariant := cleanString(input.ThemeVariant)
	if themeVariant == "" {
		themeVariant = "romantic-pink"
	}

	return s.repository.Create(
		ctx,
		title,
		message,
		openDate,
		createdBy,
		visibleTo,
		cleanOptionalString(input.CoverImageURL),
		cleanOptionalString(input.MusicURL),
		themeVariant,
	)
}

func (s *CapsuleService) Open(ctx context.Context, id int64) (*domain.LoveCapsule, error) {
	item, err := s.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if !item.IsOpened && item.OpenDate.After(time.Now()) {
		return nil, errors.Join(ErrForbidden, fmt.Errorf("capsule unlocks on %s", item.OpenDate.Format(time.RFC3339)))
	}

	if item.IsOpened {
		return item, nil
	}

	opened, err := s.repository.Open(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.Join(ErrNotFound, errors.New("capsule not found"))
		}
		return nil, err
	}
	return opened, nil
}

func (s *CapsuleService) ListScenes(ctx context.Context, capsuleID int64) ([]domain.LoveCapsuleScene, error) {
	if _, err := s.Get(ctx, capsuleID); err != nil {
		return nil, err
	}
	return s.repository.ListScenes(ctx, capsuleID)
}

func (s *CapsuleService) CreateScene(
	ctx context.Context,
	capsuleID int64,
	input domain.CreateLoveCapsuleSceneInput,
) (*domain.LoveCapsuleScene, error) {
	if _, err := s.Get(ctx, capsuleID); err != nil {
		return nil, err
	}

	if input.SceneOrder <= 0 {
		return nil, errors.Join(ErrInvalidInput, errors.New("scene_order must be greater than 0"))
	}

	sceneType := cleanString(input.SceneType)
	if sceneType == "" {
		return nil, errors.Join(ErrInvalidInput, errors.New("scene_type is required"))
	}

	return s.repository.CreateScene(
		ctx,
		capsuleID,
		input.SceneOrder,
		sceneType,
		cleanOptionalString(input.Title),
		cleanOptionalString(input.Content),
		cleanOptionalString(input.ImageURL),
		cleanOptionalString(input.AnimationKey),
	)
}
