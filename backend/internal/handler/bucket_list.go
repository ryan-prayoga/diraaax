package handler

import (
	"net/http"

	"github.com/ryanprayoga/diraaax/backend/internal/domain"
	"github.com/ryanprayoga/diraaax/backend/internal/httpresponse"
	"github.com/ryanprayoga/diraaax/backend/internal/middleware"
	"github.com/ryanprayoga/diraaax/backend/internal/service"
)

type BucketListHandler struct {
	service *service.BucketListService
}

func NewBucketListHandler(service *service.BucketListService) *BucketListHandler {
	return &BucketListHandler{service: service}
}

func (h *BucketListHandler) List(w http.ResponseWriter, r *http.Request) {
	items, err := h.service.List(r.Context())
	if err != nil {
		handleServiceError(w, err)
		return
	}
	httpresponse.Success(w, http.StatusOK, items)
}

func (h *BucketListHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input domain.CreateBucketListItemInput
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

func (h *BucketListHandler) Toggle(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(r, "id")
	if err != nil {
		httpresponse.Error(w, http.StatusBadRequest, "invalid_request", "invalid bucket list item id")
		return
	}

	authSession := middleware.CurrentAuth(r)
	item, err := h.service.Toggle(r.Context(), id, authSession.ActorUserID())
	if err != nil {
		handleServiceError(w, err)
		return
	}
	httpresponse.Success(w, http.StatusOK, item)
}

func (h *BucketListHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(r, "id")
	if err != nil {
		httpresponse.Error(w, http.StatusBadRequest, "invalid_request", "invalid bucket list item id")
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		handleServiceError(w, err)
		return
	}
	httpresponse.Success(w, http.StatusOK, map[string]int64{"deleted_id": id})
}
