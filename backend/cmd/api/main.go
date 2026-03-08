package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	"github.com/ryanprayoga/diraaax/backend/internal/config"
	"github.com/ryanprayoga/diraaax/backend/internal/db"
	"github.com/ryanprayoga/diraaax/backend/internal/handlers"
	"github.com/ryanprayoga/diraaax/backend/internal/middleware"
	"github.com/ryanprayoga/diraaax/backend/internal/repositories"
	"github.com/ryanprayoga/diraaax/backend/internal/services"
)

func main() {
	// Load .env file (ignore error if not found)
	godotenv.Load()

	cfg := config.Load()

	// Connect to database
	pool, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer pool.Close()
	log.Println("Connected to PostgreSQL")

	// Initialize repositories
	authRepo := repositories.NewAuthRepository(pool)
	planRepo := repositories.NewPlanRepository(pool)
	galleryRepo := repositories.NewGalleryRepository(pool)
	noteRepo := repositories.NewSecretNoteRepository(pool)

	// Initialize services
	authService := services.NewAuthService(authRepo)
	planService := services.NewPlanService(planRepo)
	galleryService := services.NewGalleryService(galleryRepo, cfg.UploadDir)
	noteService := services.NewSecretNoteService(noteRepo)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService, cfg.SessionCookieName)
	planHandler := handlers.NewPlanHandler(planService)
	galleryHandler := handlers.NewGalleryHandler(galleryService)
	noteHandler := handlers.NewSecretNoteHandler(noteService)

	// Setup router
	r := chi.NewRouter()

	// Global middleware
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{cfg.CORSOrigin},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Health check
	r.Get("/health", handlers.HealthCheck)

	// Serve uploaded images
	fileServer := http.StripPrefix("/uploads/", http.FileServer(http.Dir(cfg.UploadDir)))
	r.Get("/uploads/*", func(w http.ResponseWriter, r *http.Request) {
		fileServer.ServeHTTP(w, r)
	})

	// Auth routes (no auth middleware)
	r.Post("/api/auth/verify-pin", authHandler.VerifyPIN)

	// Auth routes (with auth middleware for logout and me)
	r.Group(func(r chi.Router) {
		r.Use(middleware.Auth(authService, cfg.SessionCookieName))

		r.Post("/api/auth/logout", authHandler.Logout)
		r.Get("/api/auth/me", authHandler.Me)

		// Plans
		r.Get("/api/plans", planHandler.List)
		r.Post("/api/plans", planHandler.Create)
		r.Patch("/api/plans/{id}/toggle", planHandler.ToggleStatus)
		r.Delete("/api/plans/{id}", planHandler.Delete)

		// Gallery
		r.Get("/api/gallery", galleryHandler.List)
		r.Post("/api/gallery", galleryHandler.Upload)
		r.Delete("/api/gallery/{id}", galleryHandler.Delete)

		// Secret notes
		r.Get("/api/secret-notes", noteHandler.List)
		r.Post("/api/secret-notes", noteHandler.Create)
	})

	// Start server
	server := &http.Server{
		Addr:         ":" + cfg.AppPort,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Graceful shutdown
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		log.Println("Shutting down server...")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("Server forced to shutdown: %v", err)
		}
	}()

	log.Printf("diraaax backend starting on :%s", cfg.AppPort)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server error: %v", err)
	}
	log.Println("Server stopped")
}
