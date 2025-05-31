package application

import (
	"context"
	"github.com/FindZach/TextSnapz/internal/domain"
	"github.com/matoous/go-nanoid"
	_ "strings"
	"time"
)

// TextSnapzService handles business logic for snaps.
type TextSnapzService struct {
	repo domain.TextSnapRepository
}

// NewTextSnapzService creates a new service instance.
func NewTextSnapzService(repo domain.TextSnapRepository) *TextSnapzService {
	return &TextSnapzService{repo: repo}
}

// CreateSnap creates a new snap with the given title, content, syntax, creator IP, and optional expiration.
func (s *TextSnapzService) CreateSnap(ctx context.Context, title, content, syntax, creatorIP string, expiresIn *string, tags []string) (string, error) {
	id, err := gonanoid.Generate("abcdefghijklmnopqrstuvwxyz0123456789", 10)
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
		ID:          id,
		Title:       title,
		Content:     content,
		Syntax:      syntax,
		CreatorIP:   creatorIP,
		CreatedAt:   time.Now(),
		ExpiresAt:   expiresAt,
		IsEncrypted: false,
		IsPrivate:   false,
		Tags:        tags,
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
