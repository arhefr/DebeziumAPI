package postgrespool

import "time"

type Config struct {
	User         string        `env:"POSTGRES_USER" env-default:"postgres"`
	Password     string        `env:"POSTGRES_PASSWORD" env-default:"postgres"`
	Host         string        `env:"POSTGRES_HOST" env-default:"localhost"`
	Port         string        `env:"POSTGRES_PORT" env-default:"5432"`
	Database     string        `env:"POSTGRES_DB" env-default:"postgres"`
	MaxAttemps   int           `env:"POSTGRES_ATTEMPS" env-default:"5"`
	DelayAttemps time.Duration `env:"POSTGRES_ATTEMPS_DELAY" env-default:"3s"`
}
