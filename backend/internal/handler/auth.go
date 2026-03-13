package handler

import (
	"net"
	"net/http"
	"time"

	"github.com/ryanprayoga/diraaax/backend/internal/domain"
	"github.com/ryanprayoga/diraaax/backend/internal/httpresponse"
	"github.com/ryanprayoga/diraaax/backend/internal/middleware"
	"github.com/ryanprayoga/diraaax/backend/internal/service"
)

type AuthHandler struct {
	service      *service.AuthService
	cookieName   string
	cookieSecure bool
}

func NewAuthHandler(service *service.AuthService, cookieName string, cookieSecure bool) *AuthHandler {
	return &AuthHandler{
		service:      service,
		cookieName:   cookieName,
		cookieSecure: cookieSecure,
	}
}

func (h *AuthHandler) VerifyPIN(w http.ResponseWriter, r *http.Request) {
	var input domain.VerifyPINInput
	if err := httpresponse.Decode(r, &input); err != nil {
		httpresponse.Error(w, http.StatusBadRequest, "invalid_request", "invalid JSON body")
		return
	}

	userAgent := optionalHeader(r.UserAgent())
	ipAddress := clientIP(r.RemoteAddr)

	authSession, token, err := h.service.VerifyPIN(r.Context(), input.PIN, userAgent, ipAddress)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     h.cookieName,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   h.cookieSecure,
		Expires:  authSession.Session.ExpiresAt,
		MaxAge:   int(time.Until(authSession.Session.ExpiresAt).Seconds()),
	})

	httpresponse.Success(w, http.StatusOK, authSession)
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(h.cookieName)
	if err == nil && cookie.Value != "" {
		if err := h.service.Logout(r.Context(), cookie.Value); err != nil {
			handleServiceError(w, err)
			return
		}
	}

	http.SetCookie(w, &http.Cookie{
		Name:     h.cookieName,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   h.cookieSecure,
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
	})

	httpresponse.Success(w, http.StatusOK, map[string]string{
		"message": "logged out",
	})
}

func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	authSession := middleware.CurrentAuth(r)
	if authSession == nil {
		httpresponse.Error(w, http.StatusUnauthorized, "unauthorized", "authentication required")
		return
	}

	httpresponse.Success(w, http.StatusOK, authSession)
}

func optionalHeader(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}

func clientIP(remoteAddr string) *string {
	host, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		if remoteAddr == "" {
			return nil
		}
		return &remoteAddr
	}
	return &host
}
