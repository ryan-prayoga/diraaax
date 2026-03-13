package handler

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"

	"github.com/ryanprayoga/diraaax/backend/internal/httpresponse"
	"github.com/ryanprayoga/diraaax/backend/internal/service"
)

func parseIDParam(r *http.Request, key string) (int64, error) {
	return strconv.ParseInt(chi.URLParam(r, key), 10, 64)
}

func handleServiceError(w http.ResponseWriter, err error) {
	message := serviceErrorMessage(err)

	switch {
	case errors.Is(err, service.ErrInvalidInput):
		httpresponse.Error(w, http.StatusBadRequest, "invalid_request", message)
	case errors.Is(err, service.ErrUnauthorized):
		httpresponse.Error(w, http.StatusUnauthorized, "unauthorized", message)
	case errors.Is(err, service.ErrForbidden):
		httpresponse.Error(w, http.StatusForbidden, "forbidden", message)
	case errors.Is(err, service.ErrNotFound):
		httpresponse.Error(w, http.StatusNotFound, "not_found", message)
	default:
		log.Printf("handler error: %v", err)
		httpresponse.Error(w, http.StatusInternalServerError, "internal_error", "internal server error")
	}
}

func serviceErrorMessage(err error) string {
	parts := strings.Split(err.Error(), "\n")
	for i := len(parts) - 1; i >= 0; i-- {
		if trimmed := strings.TrimSpace(parts[i]); trimmed != "" {
			return trimmed
		}
	}
	return "request failed"
}
