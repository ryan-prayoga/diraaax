package handlers

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/ryanprayoga/diraaax/backend/internal/middleware"
	"github.com/ryanprayoga/diraaax/backend/internal/services"
	"github.com/ryanprayoga/diraaax/backend/internal/utils"
)

type AuthHandler struct {
	service    *services.AuthService
	cookieName string
}

func NewAuthHandler(service *services.AuthService, cookieName string) *AuthHandler {
	return &AuthHandler{service: service, cookieName: cookieName}
}

func (h *AuthHandler) VerifyPIN(w http.ResponseWriter, r *http.Request) {
	var body struct {
		PIN string `json:"pin"`
	}
	if err := utils.DecodeJSON(r, &body); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "PIN is required")
		return
	}

	body.PIN = strings.TrimSpace(body.PIN)
	if body.PIN == "" {
		utils.RespondError(w, http.StatusBadRequest, "PIN is required")
		return
	}

	result, err := h.service.VerifyPIN(r.Context(), body.PIN)
	if err != nil {
		if errors.Is(err, services.ErrInvalidPIN) {
			log.Printf("[AUTH] PIN verification rejected from %s", r.RemoteAddr)
			utils.RespondError(w, http.StatusUnauthorized, "invalid_pin")
			return
		}

		log.Printf("[AUTH] Failed to create login session from %s: %v", r.RemoteAddr, err)
		utils.RespondError(w, http.StatusInternalServerError, "login_failed")
		return
	}

	log.Printf("[AUTH] Successful login from %s", r.RemoteAddr)

	http.SetCookie(w, &http.Cookie{
		Name:     h.cookieName,
		Value:    result.Session.Token,
		Path:     "/",
		Expires:  result.Session.ExpiresAt,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   false, // set true in production with HTTPS
	})

	utils.RespondJSON(w, http.StatusOK, map[string]interface{}{
		"message": "Login successful",
		"user":    result.User,
	})
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(h.cookieName)
	if err == nil && cookie.Value != "" {
		h.service.Logout(r.Context(), cookie.Value)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     h.cookieName,
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   -1,
	})

	utils.RespondJSON(w, http.StatusOK, map[string]string{"message": "Logged out"})
}

func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	auth := middleware.GetAuth(r)
	if auth == nil {
		utils.RespondError(w, http.StatusUnauthorized, "Not authenticated")
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]interface{}{
		"authenticated": true,
		"user":          auth.User,
	})
}
