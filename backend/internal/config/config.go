package config

import (
	"errors"
	"os"
	"strings"
	"time"
)

const defaultSessionTTL = 7 * 24 * time.Hour

type Config struct {
	AppPort           string
	DatabaseURL       string
	SessionCookieName string
	SessionSecret     string
	UploadDir         string
	CORSOrigin        string
	CookieSecure      bool
	SessionTTL        time.Duration
}

func Load() (Config, error) {
	cfg := Config{
		AppPort:           getEnv("APP_PORT", "8080"),
		DatabaseURL:       strings.TrimSpace(os.Getenv("DATABASE_URL")),
		SessionCookieName: getEnv("SESSION_COOKIE_NAME", "diraaax_session"),
		SessionSecret:     strings.TrimSpace(os.Getenv("SESSION_SECRET")),
		UploadDir:         getEnv("UPLOAD_DIR", "./uploads"),
		CORSOrigin:        strings.TrimRight(getEnv("CORS_ORIGIN", "http://localhost:5173"), "/"),
		SessionTTL:        defaultSessionTTL,
	}

	if cfg.DatabaseURL == "" {
		return Config{}, errors.New("DATABASE_URL is required")
	}
	if cfg.SessionSecret == "" {
		return Config{}, errors.New("SESSION_SECRET is required")
	}

	cfg.CookieSecure = !isLocalOrigin(cfg.CORSOrigin)
	return cfg, nil
}

func getEnv(key, fallback string) string {
	if val := strings.TrimSpace(os.Getenv(key)); val != "" {
		return val
	}
	return fallback
}

func isLocalOrigin(origin string) bool {
	return strings.Contains(origin, "localhost") || strings.Contains(origin, "127.0.0.1")
}
