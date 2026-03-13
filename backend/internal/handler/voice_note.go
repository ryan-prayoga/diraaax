package handler

import (
	"net/http"

	"github.com/ryanprayoga/diraaax/backend/internal/domain"
	"github.com/ryanprayoga/diraaax/backend/internal/httpresponse"
	"github.com/ryanprayoga/diraaax/backend/internal/middleware"
	"github.com/ryanprayoga/diraaax/backend/internal/service"
)

type VoiceNoteHandler struct {
	service *service.VoiceNoteService
}

func NewVoiceNoteHandler(service *service.VoiceNoteService) *VoiceNoteHandler {
	return &VoiceNoteHandler{service: service}
}

func (h *VoiceNoteHandler) List(w http.ResponseWriter, r *http.Request) {
	items, err := h.service.List(r.Context())
	if err != nil {
		handleServiceError(w, err)
		return
	}
	httpresponse.Success(w, http.StatusOK, items)
}

func (h *VoiceNoteHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input domain.CreateVoiceNoteInput
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
