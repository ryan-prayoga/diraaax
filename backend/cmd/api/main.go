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
	"github.com/ryanprayoga/diraaax/backend/internal/handler"
	"github.com/ryanprayoga/diraaax/backend/internal/middleware"
	"github.com/ryanprayoga/diraaax/backend/internal/repository"
	"github.com/ryanprayoga/diraaax/backend/internal/service"
)

func main() {
	_ = godotenv.Load()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	ctx := context.Background()
	pool, err := db.NewPostgresPool(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("connect postgres: %v", err)
	}
	defer pool.Close()

	authRepository := repository.NewAuthRepository(pool, cfg.SessionSecret)
	timelineRepository := repository.NewTimelineRepository(pool)
	memoryRepository := repository.NewMemoryRepository(pool)
	bucketListRepository := repository.NewBucketListRepository(pool)
	capsuleRepository := repository.NewCapsuleRepository(pool)
	moodRepository := repository.NewMoodRepository(pool)
	loveReasonRepository := repository.NewLoveReasonRepository(pool)
	voiceNoteRepository := repository.NewVoiceNoteRepository(pool)
	memoryLocationRepository := repository.NewMemoryLocationRepository(pool)

	authService := service.NewAuthService(authRepository, cfg.SessionTTL)
	timelineService := service.NewTimelineService(timelineRepository)
	memoryService := service.NewMemoryService(memoryRepository)
	bucketListService := service.NewBucketListService(bucketListRepository)
	capsuleService := service.NewCapsuleService(capsuleRepository)
	moodService := service.NewMoodService(moodRepository)
	loveReasonService := service.NewLoveReasonService(loveReasonRepository)
	voiceNoteService := service.NewVoiceNoteService(voiceNoteRepository)
	memoryLocationService := service.NewMemoryLocationService(memoryLocationRepository)

	healthHandler := handler.NewHealthHandler(pool)
	authHandler := handler.NewAuthHandler(authService, cfg.SessionCookieName, cfg.CookieSecure)
	timelineHandler := handler.NewTimelineHandler(timelineService)
	memoryHandler := handler.NewMemoryHandler(memoryService)
	bucketListHandler := handler.NewBucketListHandler(bucketListService)
	capsuleHandler := handler.NewCapsuleHandler(capsuleService)
	moodHandler := handler.NewMoodHandler(moodService)
	loveReasonHandler := handler.NewLoveReasonHandler(loveReasonService)
	voiceNoteHandler := handler.NewVoiceNoteHandler(voiceNoteService)
	memoryLocationHandler := handler.NewMemoryLocationHandler(memoryLocationService)

	router := chi.NewRouter()
	router.Use(middleware.RequestLogger)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{cfg.CORSOrigin},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Get("/", healthHandler.Get)
	router.Get("/health", healthHandler.Get)

	router.Route("/api", func(api chi.Router) {
		api.Post("/auth/verify-pin", authHandler.VerifyPIN)

		api.Group(func(private chi.Router) {
			private.Use(middleware.RequireAuth(cfg.SessionCookieName, authService))

			private.Post("/auth/logout", authHandler.Logout)
			private.Get("/auth/me", authHandler.Me)

			private.Get("/timeline", timelineHandler.List)
			private.Post("/timeline", timelineHandler.Create)
			private.Delete("/timeline/{id}", timelineHandler.Delete)

			private.Get("/memories", memoryHandler.List)
			private.Get("/memories/random", memoryHandler.Random)
			private.Post("/memories", memoryHandler.Create)
			private.Delete("/memories/{id}", memoryHandler.Delete)

			private.Get("/bucket-list", bucketListHandler.List)
			private.Post("/bucket-list", bucketListHandler.Create)
			private.Patch("/bucket-list/{id}/toggle", bucketListHandler.Toggle)
			private.Delete("/bucket-list/{id}", bucketListHandler.Delete)

			private.Get("/capsules", capsuleHandler.List)
			private.Post("/capsules", capsuleHandler.Create)
			private.Get("/capsules/{id}", capsuleHandler.Get)
			private.Post("/capsules/{id}/open", capsuleHandler.Open)
			private.Get("/capsules/{id}/scenes", capsuleHandler.ListScenes)
			private.Post("/capsules/{id}/scenes", capsuleHandler.CreateScene)

			private.Get("/moods", moodHandler.List)
			private.Post("/moods", moodHandler.Create)

			private.Get("/love-reasons", loveReasonHandler.List)
			private.Post("/love-reasons", loveReasonHandler.Create)
			private.Delete("/love-reasons/{id}", loveReasonHandler.Delete)

			private.Get("/voice-notes", voiceNoteHandler.List)
			private.Post("/voice-notes", voiceNoteHandler.Create)

			private.Get("/memory-locations", memoryLocationHandler.List)
			private.Post("/memory-locations", memoryLocationHandler.Create)
		})
	})

	server := &http.Server{
		Addr:              ":" + cfg.AppPort,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	go gracefulShutdown(server)

	log.Printf("diraaax backend listening on :%s", cfg.AppPort)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("serve http: %v", err)
	}
}

func gracefulShutdown(server *http.Server) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("shutdown error: %v", err)
	}
}
