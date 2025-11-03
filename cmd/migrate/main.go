package main

import (
	"debez/internal/config"
	"debez/pkg/postgrespool"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var migrationsPath string
	var envPath string
	var command string

	flag.StringVar(&migrationsPath, "path", "./migrations", "path to migrations")
	flag.StringVar(&command, "command", "up", "type of command")
	flag.StringVar(&envPath, "env", "./config/.env", "path to env")
	flag.Parse()

	cfg, err := config.ParseConfig(envPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dsn := postgrespool.ConnURL(cfg.PostgresConfig)

	m, err := migrate.New(
		fmt.Sprintf("file://%s", migrationsPath),
		dsn,
	)
	if err != nil {
		fmt.Printf("failed to create migrate instance: %v\n", err)
		os.Exit(1)
	}
	defer m.Close()

	switch command {

	case "up":

		if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			fmt.Printf("failed to apply migrations: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Migrations applied successfully!")

	case "down":

		if err := m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			fmt.Printf("failed to rollback migrations: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Migrations rolled back successfully!")

	case "version":

		version, dirty, err := m.Version()
		if err != nil {
			fmt.Printf("failed to get version: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Current version: %d, Dirty: %v\n", version, dirty)

	default:

		fmt.Printf("unknown command %s\n", command)
		fmt.Println("Available commands: up, down, version")
		os.Exit(1)
	}
}
