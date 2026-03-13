package service

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"github.com/ryanprayoga/diraaax/backend/internal/domain"
	"github.com/ryanprayoga/diraaax/backend/internal/repository"
)

type LoveReasonService struct {
	repository *repository.LoveReasonRepository
}

func NewLoveReasonService(repository *repository.LoveReasonRepository) *LoveReasonService {
	return &LoveReasonService{repository: repository}
}

func (s *LoveReasonService) List(ctx context.Context) ([]domain.LoveReason, error) {
	return s.repository.List(ctx)
}

func (s *LoveReasonService) Create(
	ctx context.Context,
	createdBy *int64,
	input domain.CreateLoveReasonInput,
) (*domain.LoveReason, error) {
	message := cleanString(input.Message)
	if message == "" {
		return nil, errors.Join(ErrInvalidInput, errors.New("message is required"))
	}

	visibleTo := cleanString(input.VisibleTo)
	if visibleTo == "" {
		visibleTo = "both"
	}

	return s.repository.Create(ctx, message, createdBy, visibleTo, input.IsPinned)
}

func (s *LoveReasonService) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.Join(ErrInvalidInput, errors.New("invalid love reason id"))
	}
	if err := s.repository.Delete(ctx, id); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errors.Join(ErrNotFound, errors.New("love reason not found"))
		}
		return err
	}
	return nil
}
