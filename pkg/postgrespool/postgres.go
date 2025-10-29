package postgrespool

import (
	"context"
	"fmt"
	"time"

	"debez/pkg/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrConnPostgres = "error cannot connecting to postgres pool with `postgres://%s:%s@%s:%s/%s`"
)

const (
	defaultConnTimeout = 5 * time.Second
)

func NewPool(ctx context.Context, cfg *Config) (pool *pgxpool.Pool, err error) {

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	err = utils.DoWithTries(func() error {

		ctx, cancel := context.WithTimeout(ctx, defaultConnTimeout)
		defer cancel()

		pool, err = pgxpool.New(ctx, connStr)
		if err != nil {
			return fmt.Errorf("error failed connect to postgres pool")
		}

		return nil
	}, cfg.MaxAttemps, cfg.DelayAttemps)

	if err != nil {
		return nil, fmt.Errorf(ErrConnPostgres, cfg.User, "<password>", cfg.Host, cfg.Port, cfg.Database)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf(ErrConnPostgres, cfg.User, "<password>", cfg.Host, cfg.Port, cfg.Database)
	}

	return pool, nil
}
