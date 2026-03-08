package handlers

import (
	"net/http"

	"github.com/ryanprayoga/diraaax/backend/internal/utils"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	utils.RespondJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
		"app":    "diraaax",
	})
}
