package sqlite

import (
	"context"
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DSN            = "./test_db.sqlite"
	migrationsPath = "./migrations"
	testPort       = "1234"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	ctx := context.Background()

	db, err := SetupTestDatabase(ctx)
	if err != nil {
		log.Printf("failed setup test database: %s\n", err)
		return
	}
	defer DeleteDatabase(ctx)
	defer db.Close()

	testDB = db

	m.Run()
}

func SetupTestDatabase(ctx context.Context) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", DSN)
	if err != nil {
		return nil, err
	}

	var query string

	migrations, err := filepath.Glob(filepath.Join(migrationsPath, "*up.sql"))
	if err != nil {
		return nil, err
	}

	for _, mig := range migrations {
		migContent, err := os.ReadFile(mig)
		if err != nil {
			return nil, err
		}
		query += string(migContent)
	}

	_, err = db.ExecContext(ctx, query)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func DeleteDatabase(ctx context.Context) error {
	if err := os.Remove(DSN); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
