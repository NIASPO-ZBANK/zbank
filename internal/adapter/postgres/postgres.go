package postgres

import (
	"fmt"
	"os"

	"users/internal/infrastructure/database"
	"users/internal/infrastructure/database/postgres"
)

func Connect(cfg *database.Config) (*postgres.Postgres, error) {
	user := os.Getenv(cfg.UserEnvKey)
	if user == "" {
		return nil, fmt.Errorf("environment variable %s is not set", cfg.UserEnvKey)
	}

	password := os.Getenv(cfg.PasswordEnvKey)
	if password == "" {
		return nil, fmt.Errorf("environment variable %s is not set", cfg.PasswordEnvKey)
	}

	url := toUrl(user, password, cfg.Addr, cfg.Port, cfg.DB)

	conn, err := postgres.New(cfg.Driver, url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	return conn, err
}

func toUrl(user, password, host, port, dbName string) string {
	return fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		user, password, host, port, dbName,
	)
}
