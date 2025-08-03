package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func NewStorageSQLite(storagePath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Migrate(storagePath string, migrationsPath string, migrationsTable string) error {
	if storagePath == "" {
		return errors.New("storage path is empty")
	}
	if migrationsPath == "" {
		return errors.New("migrations path is empty")
	}

	m, err := migrate.New(
		migrationsPath,
		fmt.Sprintf("sqlite3://%s?x-migrations-table=%s", storagePath, migrationsTable),
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations to apply")
			return nil
		}
		return err
	}

	return nil
}
