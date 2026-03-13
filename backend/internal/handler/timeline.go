package handler

import (
	"net/http"

	"github.com/ryanprayoga/diraaax/backend/internal/domain"
	"github.com/ryanprayoga/diraaax/backend/internal/httpresponse"
	"github.com/ryanprayoga/diraaax/backend/internal/middleware"
	"github.com/ryanprayoga/diraaax/backend/internal/service"
)

type TimelineHandler struct {
	service *service.TimelineService
}

func NewTimelineHandler(service *service.TimelineService) *TimelineHandler {
	return &TimelineHandler{service: service}
}

func (h *TimelineHandler) List(w http.ResponseWriter, r *http.Request) {
	items, err := h.service.List(r.Context())
	if err != nil {
		handleServiceError(w, err)
		return
	}
	httpresponse.Success(w, http.StatusOK, items)
}

func (h *TimelineHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input domain.CreateTimelineEventInput
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

func (h *TimelineHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(r, "id")
	if err != nil {
		httpresponse.Error(w, http.StatusBadRequest, "invalid_request", "invalid timeline id")
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		handleServiceError(w, err)
		return
	}

	httpresponse.Success(w, http.StatusOK, map[string]int64{"deleted_id": id})
}
