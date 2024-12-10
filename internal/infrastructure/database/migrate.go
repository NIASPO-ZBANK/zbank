package database

import (
	"embed"
	"fmt"

	"github.com/pressly/goose/v3"

	"users/internal/infrastructure/database/postgres"
)

var (
	//go:embed migrations/zbank/*.sql
	migrations embed.FS
)

func MigrateZBankDB(db *postgres.Postgres) error {
	if err := migrate(db, "migrations/zbank"); err != nil {
		return fmt.Errorf("migrate: %v", err)
	}

	return nil
}

func migrate(db *postgres.Postgres, dir string) error {
	goose.SetBaseFS(migrations)

	if err := goose.Up(db.SqlDB(), dir); err != nil {
		return fmt.Errorf("goose up: %v", err)
	}

	return nil
}
