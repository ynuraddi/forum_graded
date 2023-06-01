package sqlite

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"
	"strings"
	"time"

	"graded/config"
)

func OpenDB(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", cfg.DATABASE.DSN)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	query := ""
	migrations, err := filepath.Glob(filepath.Join("./repository/sqlite/migrations", "*.up.sql"))
	if err != nil {
		return nil, err
	}
	for _, migration := range migrations {
		tmp, err := os.ReadFile(migration)
		if err != nil {
			return nil, err
		}
		query += string(tmp)
	}

	if _, err := db.ExecContext(ctx, query); err != nil {
		return nil, err
	}

	return db, nil
}

func isDuplicate(err error) bool {
	return strings.Contains(err.Error(), "unique")
}
