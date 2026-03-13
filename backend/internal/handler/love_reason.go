package handler

import (
	"net/http"

	"github.com/ryanprayoga/diraaax/backend/internal/domain"
	"github.com/ryanprayoga/diraaax/backend/internal/httpresponse"
	"github.com/ryanprayoga/diraaax/backend/internal/middleware"
	"github.com/ryanprayoga/diraaax/backend/internal/service"
)

type LoveReasonHandler struct {
	service *service.LoveReasonService
}

func NewLoveReasonHandler(service *service.LoveReasonService) *LoveReasonHandler {
	return &LoveReasonHandler{service: service}
}

func (h *LoveReasonHandler) List(w http.ResponseWriter, r *http.Request) {
	items, err := h.service.List(r.Context())
	if err != nil {
		handleServiceError(w, err)
		return
	}
	httpresponse.Success(w, http.StatusOK, items)
}

func (h *LoveReasonHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input domain.CreateLoveReasonInput
	if err := httpresponse.Decode(r, &input); err != nil {
		httpresponse.Error(w, http.StatusBadRequest, "invalid_request", "invalid JSON body")
		return
	}

	authSession := middleware.CurrentAuth(r)
	item, err := h.service.Create(r.Context(), authSession.ActorUserID(), input)
	if err != nil {
		handleServiceError(w, err)
		return
	}
	httpresponse.Success(w, http.StatusCreated, item)
}

func (h *LoveReasonHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(r, "id")
	if err != nil {
		httpresponse.Error(w, http.StatusBadRequest, "invalid_request", "invalid love reason id")
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		handleServiceError(w, err)
		return
	}
	httpresponse.Success(w, http.StatusOK, map[string]int64{"deleted_id": id})
}
