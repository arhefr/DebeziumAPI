package postgrespool

import (
	"context"
	"fmt"
	"time"

	"debez/pkg/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrConnPostgres = "error cannot connecting to postgres pool with `postgres://%s:%s@%s:%s/%s?sslmode=disable`"
)

const (
	defaultConnTimeout = 5 * time.Second
)

type Database struct {
	Pool *pgxpool.Pool
}

func New(ctx context.Context, cfg *Config) (db *Database, err error) {
	db = &Database{}

	connStr := ConnURL(cfg)

	err = utils.DoWithTries(func() error {

		ctx, cancel := context.WithTimeout(ctx, defaultConnTimeout)
		defer cancel()

		db.Pool, err = pgxpool.New(ctx, connStr)
		if err != nil {
			return fmt.Errorf("error failed connect to postgres pool")
		}

		return nil
	}, cfg.MaxAttemps, cfg.DelayAttemps)

	if err != nil {
		return nil, fmt.Errorf(ErrConnPostgres, cfg.User, "<password>", cfg.Host, cfg.Port, cfg.Database)
	}

	if err := db.Pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf(ErrConnPostgres, cfg.User, "<password>", cfg.Host, cfg.Port, cfg.Database)
	}

	return db, nil
}

func (db *Database) Stop() {
	db.Pool.Close()
}

func ConnURL(cfg *Config) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
}
