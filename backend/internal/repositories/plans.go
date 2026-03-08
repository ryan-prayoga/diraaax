package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ryanprayoga/diraaax/backend/internal/models"
)

type PlanRepository struct {
	pool *pgxpool.Pool
}

func NewPlanRepository(pool *pgxpool.Pool) *PlanRepository {
	return &PlanRepository{pool: pool}
}

func (r *PlanRepository) List(ctx context.Context) ([]models.FuturePlan, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, title, description, category, status, created_by, completed_at, created_at
		 FROM future_plans
		 ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var plans []models.FuturePlan
	for rows.Next() {
		var p models.FuturePlan
		if err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.Category, &p.Status, &p.CreatedBy, &p.CompletedAt, &p.CreatedAt); err != nil {
			return nil, err
		}
		plans = append(plans, p)
	}
	return plans, rows.Err()
}

func (r *PlanRepository) Create(ctx context.Context, title string, description *string, category *string, createdBy *int) (*models.FuturePlan, error) {
	var p models.FuturePlan
	err := r.pool.QueryRow(ctx,
		`INSERT INTO future_plans (title, description, category, status, created_by, created_at)
		 VALUES ($1, $2, $3, 'pending', $4, NOW())
		 RETURNING id, title, description, category, status, created_by, completed_at, created_at`,
		title, description, category, createdBy,
	).Scan(&p.ID, &p.Title, &p.Description, &p.Category, &p.Status, &p.CreatedBy, &p.CompletedAt, &p.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *PlanRepository) ToggleStatus(ctx context.Context, id int) (*models.FuturePlan, error) {
	var p models.FuturePlan
	err := r.pool.QueryRow(ctx,
		`UPDATE future_plans
		 SET status = CASE WHEN status = 'pending' THEN 'done' ELSE 'pending' END,
		     completed_at = CASE WHEN status = 'pending' THEN NOW() ELSE NULL END
		 WHERE id = $1
		 RETURNING id, title, description, category, status, created_by, completed_at, created_at`,
		id,
	).Scan(&p.ID, &p.Title, &p.Description, &p.Category, &p.Status, &p.CreatedBy, &p.CompletedAt, &p.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *PlanRepository) Delete(ctx context.Context, id int) error {
	_, err := r.pool.Exec(ctx, `DELETE FROM future_plans WHERE id = $1`, id)
	return err
}
