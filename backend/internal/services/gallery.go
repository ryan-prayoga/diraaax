package services

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/ryanprayoga/diraaax/backend/internal/models"
	"github.com/ryanprayoga/diraaax/backend/internal/repositories"
	"github.com/ryanprayoga/diraaax/backend/internal/utils"
)

type GalleryService struct {
	repo      *repositories.GalleryRepository
	uploadDir string
}

func NewGalleryService(repo *repositories.GalleryRepository, uploadDir string) *GalleryService {
	return &GalleryService{repo: repo, uploadDir: uploadDir}
}

func (s *GalleryService) List(ctx context.Context) ([]models.GalleryItem, error) {
	items, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	if items == nil {
		items = []models.GalleryItem{}
	}
	return items, nil
}

type UploadInput struct {
	Title      string
	Caption    *string
	Category   *string
	TakenAt    *time.Time
	UploadedBy *int
	FileData   io.Reader
	FileName   string
}

func (s *GalleryService) Upload(ctx context.Context, input UploadInput) (*models.GalleryItem, error) {
	if err := os.MkdirAll(s.uploadDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create upload directory: %w", err)
	}

	// Generate unique filename
	token, err := utils.GenerateToken(8)
	if err != nil {
		return nil, fmt.Errorf("failed to generate filename: %w", err)
	}

	ext := filepath.Ext(input.FileName)
	uniqueName := fmt.Sprintf("%d_%s%s", time.Now().Unix(), token, ext)
	fullPath := filepath.Join(s.uploadDir, uniqueName)

	dst, err := os.Create(fullPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create file: %w", err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, input.FileData); err != nil {
		os.Remove(fullPath)
		return nil, fmt.Errorf("failed to save file: %w", err)
	}

	item, err := s.repo.Create(ctx, input.Title, input.Caption, input.Category, uniqueName, fullPath, input.TakenAt, input.UploadedBy)
	if err != nil {
		os.Remove(fullPath)
		return nil, fmt.Errorf("failed to save metadata: %w", err)
	}

	return item, nil
}

func (s *GalleryService) Delete(ctx context.Context, id int) error {
	item, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("gallery item not found")
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}

	// Best-effort file deletion
	os.Remove(item.FilePath)
	return nil
}
