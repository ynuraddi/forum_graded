package sqlite

import (
	"context"
	"database/sql"
	"os"
	"time"
)

func NewDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./forum.sqlite")
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err := os.ReadFile("./configDB.sql")
	if err != nil {
		return nil, err
	}

	if _, err := db.ExecContext(ctx, string(stmt)); err != nil {
		return nil, err
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
