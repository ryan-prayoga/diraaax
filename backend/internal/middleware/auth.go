package middleware

import (
	"context"
	"net/http"

	"github.com/ryanprayoga/diraaax/backend/internal/domain"
	"github.com/ryanprayoga/diraaax/backend/internal/httpresponse"
)

type authContextKey string

const currentAuthKey authContextKey = "current_auth"

type SessionReader interface {
	GetSession(ctx context.Context, token string) (*domain.AuthSession, error)
}

func RequireAuth(cookieName string, sessionReader SessionReader) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie(cookieName)
			if err != nil || cookie.Value == "" {
				httpresponse.Error(w, http.StatusUnauthorized, "unauthorized", "authentication required")
				return
			}

			authSession, err := sessionReader.GetSession(r.Context(), cookie.Value)
			if err != nil {
				httpresponse.Error(w, http.StatusUnauthorized, "unauthorized", "invalid or expired session")
				return
			}

			ctx := context.WithValue(r.Context(), currentAuthKey, authSession)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func CurrentAuth(r *http.Request) *domain.AuthSession {
	authSession, ok := r.Context().Value(currentAuthKey).(*domain.AuthSession)
	if !ok {
		return nil
	}
	return authSession
}
