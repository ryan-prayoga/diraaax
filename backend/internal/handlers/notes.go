package handlers

import (
	"net/http"

	"github.com/ryanprayoga/diraaax/backend/internal/middleware"
	"github.com/ryanprayoga/diraaax/backend/internal/services"
	"github.com/ryanprayoga/diraaax/backend/internal/utils"
)

type SecretNoteHandler struct {
	service *services.SecretNoteService
}

func NewSecretNoteHandler(service *services.SecretNoteService) *SecretNoteHandler {
	return &SecretNoteHandler{service: service}
}

func (h *SecretNoteHandler) List(w http.ResponseWriter, r *http.Request) {
	notes, err := h.service.List(r.Context())
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to fetch notes")
		return
	}
	utils.RespondJSON(w, http.StatusOK, notes)
}

func (h *SecretNoteHandler) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Title     string  `json:"title"`
		Content   string  `json:"content"`
		NoteType  *string `json:"note_type"`
		VisibleTo *string `json:"visible_to"`
	}
	if err := utils.DecodeJSON(r, &body); err != nil || body.Title == "" || body.Content == "" {
		utils.RespondError(w, http.StatusBadRequest, "Title and content are required")
		return
	}

	var createdBy *int
	if auth := middleware.GetAuth(r); auth != nil && auth.User != nil {
		createdBy = &auth.User.ID
	}

	note, err := h.service.Create(r.Context(), body.Title, body.Content, body.NoteType, body.VisibleTo, createdBy)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to create note")
		return
	}
	utils.RespondJSON(w, http.StatusCreated, note)
}
