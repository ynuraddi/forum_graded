package sqlite

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"graded/config"
	"graded/utils"

	_ "github.com/mattn/go-sqlite3"
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

	// TODO added migrationsPath to config
	migrations, err := utils.ParseDirectoryFiles("", "*up.sql")
	if err != nil {
		return nil, err
	}

	if _, err := db.ExecContext(ctx, migrations); err != nil {
		return nil, err
	}

	return db, nil
}

func isDuplicate(err error) bool {
	return strings.Contains(err.Error(), "UNIQUE")
}
