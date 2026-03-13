package handler

import (
	"net/http"

	"github.com/ryanprayoga/diraaax/backend/internal/domain"
	"github.com/ryanprayoga/diraaax/backend/internal/httpresponse"
	"github.com/ryanprayoga/diraaax/backend/internal/middleware"
	"github.com/ryanprayoga/diraaax/backend/internal/service"
)

type CapsuleHandler struct {
	service *service.CapsuleService
}

func NewCapsuleHandler(service *service.CapsuleService) *CapsuleHandler {
	return &CapsuleHandler{service: service}
}

func (h *CapsuleHandler) List(w http.ResponseWriter, r *http.Request) {
	items, err := h.service.List(r.Context())
	if err != nil {
		handleServiceError(w, err)
		return
	}
	httpresponse.Success(w, http.StatusOK, items)
}

func (h *CapsuleHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(r, "id")
	if err != nil {
		httpresponse.Error(w, http.StatusBadRequest, "invalid_request", "invalid capsule id")
		return
	}

	item, err := h.service.Get(r.Context(), id)
	if err != nil {
		handleServiceError(w, err)
		return
	}
	httpresponse.Success(w, http.StatusOK, item)
}

func (h *CapsuleHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input domain.CreateLoveCapsuleInput
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

func (h *CapsuleHandler) Open(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(r, "id")
	if err != nil {
		httpresponse.Error(w, http.StatusBadRequest, "invalid_request", "invalid capsule id")
		return
	}

	item, err := h.service.Open(r.Context(), id)
	if err != nil {
		handleServiceError(w, err)
		return
	}
	httpresponse.Success(w, http.StatusOK, item)
}

func (h *CapsuleHandler) ListScenes(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(r, "id")
	if err != nil {
		httpresponse.Error(w, http.StatusBadRequest, "invalid_request", "invalid capsule id")
		return
	}

	items, err := h.service.ListScenes(r.Context(), id)
	if err != nil {
		handleServiceError(w, err)
		return
	}
	httpresponse.Success(w, http.StatusOK, items)
}

func (h *CapsuleHandler) CreateScene(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(r, "id")
	if err != nil {
		httpresponse.Error(w, http.StatusBadRequest, "invalid_request", "invalid capsule id")
		return
	}

	var input domain.CreateLoveCapsuleSceneInput
	if err := httpresponse.Decode(r, &input); err != nil {
		httpresponse.Error(w, http.StatusBadRequest, "invalid_request", "invalid JSON body")
		return
	}

	item, err := h.service.CreateScene(r.Context(), id, input)
	if err != nil {
		handleServiceError(w, err)
		return
	}
	httpresponse.Success(w, http.StatusCreated, item)
}
