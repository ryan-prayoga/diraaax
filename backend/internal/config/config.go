package config

import (
	"os"
)

type Config struct {
	AppPort           string
	DatabaseURL       string
	SessionCookieName string
	SessionSecret     string
	UploadDir         string
	CORSOrigin        string
}

func Load() *Config {
	return &Config{
		AppPort:           getEnv("APP_PORT", "8080"),
		DatabaseURL:       getEnv("DATABASE_URL", ""),
		SessionCookieName: getEnv("SESSION_COOKIE_NAME", "diraaax_session"),
		SessionSecret:     getEnv("SESSION_SECRET", "change_me_in_production"),
		UploadDir:         getEnv("UPLOAD_DIR", "./uploads/gallery"),
		CORSOrigin:        getEnv("CORS_ORIGIN", "http://localhost:5173"),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
