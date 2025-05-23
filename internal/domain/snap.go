package domain

import (
	"context"
	"time"
)

// TextSnap represents a text snap (paste) in the system.
type TextSnap struct {
	ID          string
	Content     string
	CreatedAt   time.Time
	ExpiresAt   *time.Time // Nil if no expiration
	IsEncrypted bool
}

// TextSnapzRepository defines the interface for snap storage operations.
type TextSnapzRepository interface {
	Save(ctx context.Context, snap TextSnap) error
	FindByID(ctx context.Context, id string) (*TextSnap, error)
	Delete(ctx context.Context, id string) error
}
