package domain

import (
	"context"
	"time"
)

// TextSnap represents a text snap (paste) in the system.
type TextSnap struct {
	ID          string     // Unique slug (e.g., "k7b9z3m2p8")
	Title       string     // Optional title
	Content     string     // Snap content
	Syntax      string     // Syntax highlighting language (e.g., "plaintext", "javascript")
	CreatorIP   string     // IP address of the snap creator
	CreatedAt   time.Time  // Creation timestamp
	ExpiresAt   *time.Time // Optional expiration timestamp
	IsEncrypted bool       // Whether content is encrypted
	IsPrivate   bool       // Whether snap is private (future access control)
	Tags        []string   // Optional tags for categorization
}

// TextSnapRepository defines the interface for snap storage operations.
type TextSnapRepository interface {
	Save(ctx context.Context, snap TextSnap) error
	FindByID(ctx context.Context, id string) (*TextSnap, error)
	Delete(ctx context.Context, id string) error
}
