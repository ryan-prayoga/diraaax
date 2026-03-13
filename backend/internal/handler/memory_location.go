package handler

import (
	"net/http"

	"github.com/ryanprayoga/diraaax/backend/internal/domain"
	"github.com/ryanprayoga/diraaax/backend/internal/httpresponse"
	"github.com/ryanprayoga/diraaax/backend/internal/service"
)

type MemoryLocationHandler struct {
	service *service.MemoryLocationService
}

func NewMemoryLocationHandler(service *service.MemoryLocationService) *MemoryLocationHandler {
	return &MemoryLocationHandler{service: service}
}

func (h *MemoryLocationHandler) List(w http.ResponseWriter, r *http.Request) {
	items, err := h.service.List(r.Context())
	if err != nil {
		handleServiceError(w, err)
		return
	}
	httpresponse.Success(w, http.StatusOK, items)
}

func (h *MemoryLocationHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input domain.CreateMemoryLocationInput
	if err := httpresponse.Decode(r, &input); err != nil {
		httpresponse.Error(w, http.StatusBadRequest, "invalid_request", "invalid JSON body")
		return
	}

	item, err := h.service.Create(r.Context(), input)
	if err != nil {
		handleServiceError(w, err)
		return
	}
	httpresponse.Success(w, http.StatusCreated, item)
}
