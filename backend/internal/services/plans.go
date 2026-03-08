package services

import (
	"context"

	"github.com/ryanprayoga/diraaax/backend/internal/models"
	"github.com/ryanprayoga/diraaax/backend/internal/repositories"
)

type PlanService struct {
	repo *repositories.PlanRepository
}

func NewPlanService(repo *repositories.PlanRepository) *PlanService {
	return &PlanService{repo: repo}
}

func (s *PlanService) List(ctx context.Context) ([]models.FuturePlan, error) {
	plans, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	if plans == nil {
		plans = []models.FuturePlan{}
	}
	return plans, nil
}

func (s *PlanService) Create(ctx context.Context, title string, description *string, category *string, createdBy *int) (*models.FuturePlan, error) {
	return s.repo.Create(ctx, title, description, category, createdBy)
}

func (s *PlanService) ToggleStatus(ctx context.Context, id int) (*models.FuturePlan, error) {
	return s.repo.ToggleStatus(ctx, id)
}

func (s *PlanService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
