package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/ryanprayoga/diraaax/backend/internal/httpresponse"
)

type HealthHandler struct {
	pool *pgxpool.Pool
}

func NewHealthHandler(pool *pgxpool.Pool) *HealthHandler {
	return &HealthHandler{pool: pool}
}

func (h *HealthHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	if err := h.pool.Ping(ctx); err != nil {
		httpresponse.Error(w, http.StatusServiceUnavailable, "service_unavailable", "database unavailable")
		return
	}

	httpresponse.Success(w, http.StatusOK, map[string]any{
		"status": "ok",
	})
}
