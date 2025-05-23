package application

import (
	"context"
	"github.com/FindZach/TextSnapz/internal/domain"
	"github.com/matoous/go-nanoid/v2"
	"time"
)

// TextSnapzService handles business logic for snaps.
type TextSnapzService struct {
	repo domain.TextSnapzRepository
}

// NewTextSnapzService creates a new service instance.
func NewTextSnapzService(repo domain.TextSnapzRepository) *TextSnapzService {
	return &TextSnapzService{repo: repo}
}

// CreateSnap creates a new snap with the given content and optional expiration.
func (s *TextSnapzService) CreateSnap(ctx context.Context, content string, expiresIn *string) (string, error) {
	id, err := gonanoid.New()
	if err != nil {
		return "", err
	}

	var expiresAt *time.Time
	if expiresIn != nil {
		duration, err := time.ParseDuration(*expiresIn)
		if err != nil {
			return "", err
		}
		t := time.Now().Add(duration)
		expiresAt = &t
	}

	snap := domain.TextSnap{
		ID:        id,
		Content:   content,
		CreatedAt: time.Now(),
		ExpiresAt: expiresAt,
	}

	if err := s.repo.Save(ctx, snap); err != nil {
		return "", err
	}
	return id, nil
}

// GetSnap retrieves a snap by ID, checking for expiration.
func (s *TextSnapzService) GetSnap(ctx context.Context, id string) (*domain.TextSnap, error) {
	snap, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if snap.ExpiresAt != nil && snap.ExpiresAt.Before(time.Now()) {
		return nil, domain.ErrSnapExpired
	}
	return snap, nil
}
