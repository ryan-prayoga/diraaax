package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/ryanprayoga/diraaax/backend/internal/services"
	"github.com/ryanprayoga/diraaax/backend/internal/utils"
)

type GalleryHandler struct {
	service *services.GalleryService
}

func NewGalleryHandler(service *services.GalleryService) *GalleryHandler {
	return &GalleryHandler{service: service}
}

func (h *GalleryHandler) List(w http.ResponseWriter, r *http.Request) {
	items, err := h.service.List(r.Context())
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to fetch gallery")
		return
	}
	utils.RespondJSON(w, http.StatusOK, items)
}

func (h *GalleryHandler) Upload(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(32 << 20); err != nil { // 32MB max
		utils.RespondError(w, http.StatusBadRequest, "Failed to parse form")
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Image file is required")
		return
	}
	defer file.Close()

	title := r.FormValue("title")
	if title == "" {
		title = header.Filename
	}

	var caption *string
	if c := r.FormValue("caption"); c != "" {
		caption = &c
	}

	var category *string
	if c := r.FormValue("category"); c != "" {
		category = &c
	}

	var takenAt *time.Time
	if t := r.FormValue("taken_at"); t != "" {
		parsed, err := time.Parse("2006-01-02", t)
		if err == nil {
			takenAt = &parsed
		}
	}

	var uploadedBy *int
	if u := r.FormValue("uploaded_by"); u != "" {
		id, err := strconv.Atoi(u)
		if err == nil {
			uploadedBy = &id
		}
	}

	input := services.UploadInput{
		Title:      title,
		Caption:    caption,
		Category:   category,
		TakenAt:    takenAt,
		UploadedBy: uploadedBy,
		FileData:   file,
		FileName:   header.Filename,
	}

	item, err := h.service.Upload(r.Context(), input)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to upload image")
		return
	}
	utils.RespondJSON(w, http.StatusCreated, item)
}

func (h *GalleryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid gallery item ID")
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to delete gallery item")
		return
	}
	utils.RespondJSON(w, http.StatusOK, map[string]string{"message": "Gallery item deleted"})
}
