package service

import (
	"context"
	"errors"

	"github.com/ryanprayoga/diraaax/backend/internal/domain"
	"github.com/ryanprayoga/diraaax/backend/internal/repository"
)

type VoiceNoteService struct {
	repository *repository.VoiceNoteRepository
}

func NewVoiceNoteService(repository *repository.VoiceNoteRepository) *VoiceNoteService {
	return &VoiceNoteService{repository: repository}
}

func (s *VoiceNoteService) List(ctx context.Context) ([]domain.VoiceNote, error) {
	return s.repository.List(ctx)
}

func (s *VoiceNoteService) Create(
	ctx context.Context,
	createdBy *int64,
	input domain.CreateVoiceNoteInput,
) (*domain.VoiceNote, error) {
	audioURL := cleanString(input.AudioURL)
	if audioURL == "" {
		return nil, errors.Join(ErrInvalidInput, errors.New("audio_url is required"))
	}

	return s.repository.Create(
		ctx,
		cleanOptionalString(input.Title),
		audioURL,
		input.DurationSeconds,
		cleanOptionalString(input.Transcript),
		createdBy,
	)
}
