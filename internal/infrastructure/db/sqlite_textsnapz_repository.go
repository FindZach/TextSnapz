package db

import (
	"context"
	"database/sql"
	"github.com/FindZach/TextSnapz/internal/domain"
	_ "github.com/mattn/go-sqlite3"
	"strings"
)

// SqliteTextSnapRepository implements TextSnapRepository using SQLite.
type SqliteTextSnapRepository struct {
	db *sql.DB
}

// NewSqliteTextSnapRepository creates a new repository instance.
func NewSqliteTextSnapRepository(db *sql.DB) *SqliteTextSnapRepository {
	return &SqliteTextSnapRepository{db: db}
}

// Save stores a snap in the database.
func (r *SqliteTextSnapRepository) Save(ctx context.Context, snap domain.TextSnap) error {
	query := `INSERT INTO snaps (id, title, content, syntax, creator_ip, created_at, expires_at, is_encrypted, is_private, tags)
              VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	tags := strings.Join(snap.Tags, ",")
	_, err := r.db.ExecContext(ctx, query,
		snap.ID, snap.Title, snap.Content, snap.Syntax, snap.CreatorIP, snap.CreatedAt, snap.ExpiresAt, snap.IsEncrypted, snap.IsPrivate, tags)
	return err
}

// FindByID retrieves a snap by its ID.
func (r *SqliteTextSnapRepository) FindByID(ctx context.Context, id string) (*domain.TextSnap, error) {
	query := `SELECT id, title, content, syntax, creator_ip, created_at, expires_at, is_encrypted, is_private, tags FROM snaps WHERE id = ?`
	row := r.db.QueryRowContext(ctx, query, id)
	var snap domain.TextSnap
	var expiresAt sql.NullTime
	var tagsStr string
	err := row.Scan(&snap.ID, &snap.Title, &snap.Content, &snap.Syntax, &snap.CreatorIP, &snap.CreatedAt, &expiresAt, &snap.IsEncrypted, &snap.IsPrivate, &tagsStr)
	if err == sql.ErrNoRows {
		return nil, domain.ErrSnapNotFound
	}
	if err != nil {
		return nil, err
	}
	if expiresAt.Valid {
		snap.ExpiresAt = &expiresAt.Time
	}
	if tagsStr != "" {
		snap.Tags = strings.Split(tagsStr, ",")
	}
	return &snap, nil
}

// Delete removes a snap by its ID.
func (r *SqliteTextSnapRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM snaps WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
