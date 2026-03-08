package middleware

import (
	"context"
	"net/http"

	"github.com/ryanprayoga/diraaax/backend/internal/services"
)

type contextKey string

const AuthContextKey contextKey = "auth"

func Auth(authService *services.AuthService, cookieName string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie(cookieName)
			if err != nil || cookie.Value == "" {
				http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
				return
			}

			result, err := authService.GetSession(r.Context(), cookie.Value)
			if err != nil {
				http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), AuthContextKey, result)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetAuth(r *http.Request) *services.AuthResult {
	result, ok := r.Context().Value(AuthContextKey).(*services.AuthResult)
	if !ok {
		return nil
	}
	return result
}
