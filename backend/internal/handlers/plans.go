package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/ryanprayoga/diraaax/backend/internal/middleware"
	"github.com/ryanprayoga/diraaax/backend/internal/services"
	"github.com/ryanprayoga/diraaax/backend/internal/utils"
)

type PlanHandler struct {
	service *services.PlanService
}

func NewPlanHandler(service *services.PlanService) *PlanHandler {
	return &PlanHandler{service: service}
}

func (h *PlanHandler) List(w http.ResponseWriter, r *http.Request) {
	plans, err := h.service.List(r.Context())
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to fetch plans")
		return
	}
	utils.RespondJSON(w, http.StatusOK, plans)
}

func (h *PlanHandler) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Title       string  `json:"title"`
		Description *string `json:"description"`
		Category    *string `json:"category"`
	}
	if err := utils.DecodeJSON(r, &body); err != nil || body.Title == "" {
		utils.RespondError(w, http.StatusBadRequest, "Title is required")
		return
	}

	var createdBy *int
	if auth := middleware.GetAuth(r); auth != nil && auth.User != nil {
		createdBy = &auth.User.ID
	}

	plan, err := h.service.Create(r.Context(), body.Title, body.Description, body.Category, createdBy)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to create plan")
		return
	}
	utils.RespondJSON(w, http.StatusCreated, plan)
}

func (h *PlanHandler) ToggleStatus(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid plan ID")
		return
	}

	plan, err := h.service.ToggleStatus(r.Context(), id)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to toggle plan status")
		return
	}
	utils.RespondJSON(w, http.StatusOK, plan)
}

func (h *PlanHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid plan ID")
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to delete plan")
		return
	}
	utils.RespondJSON(w, http.StatusOK, map[string]string{"message": "Plan deleted"})
}
