package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// SetupDatabase initializes the SQLite database and creates the snaps table.
func SetupDatabase(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	createTableQuery := `
        CREATE TABLE IF NOT EXISTS snaps (
            id TEXT PRIMARY KEY,
            title TEXT,
            content TEXT NOT NULL,
            syntax TEXT,
            creator_ip TEXT,
            created_at DATETIME NOT NULL,
            expires_at DATETIME,
            is_encrypted BOOLEAN NOT NULL,
            is_private BOOLEAN NOT NULL,
            tags TEXT
        );`
	if _, err := db.Exec(createTableQuery); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
