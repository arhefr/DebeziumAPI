package postgrespool

import (
	"context"
	"errors"
	"fmt"
	"time"

	"debez/pkg/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPool(ctx context.Context, cfg *Config) (pool *pgxpool.Pool, err error) {

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	err = utils.DoWithTries(func() error {

		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgxpool.New(ctx, connStr)
		if err != nil {
			return fmt.Errorf("error failed connect to postgres pool")
		}

		return nil
	}, cfg.MaxAttemps, cfg.DelayAttemps)

	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("cannot connecting to postgres pool with `postgres://%s:%s@%s:%s/%s`",
			cfg.User, "<password>", cfg.Host, cfg.Port, cfg.Database))
	}

	return pool, nil
}
