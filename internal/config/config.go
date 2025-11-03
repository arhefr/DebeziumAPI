package config

import (
	"fmt"
	"time"

	"debez/pkg/postgrespool"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Environment string `env:"ENV" env-default:"dev"`

	PostgresConfig   *postgrespool.Config
	GRPCServerConfig *GRPCServerConfig
	HTTPServerConfig *HTTPServerConfig
}

type HTTPServerConfig struct {
	Port            string        `env:"HTTP_PORT"         env-default:"8080"`
	Timeout         time.Duration `env:"HTTP_TIMEOUT" env-default:"30s"`
	DebeziumBaseURL string        `env:"DEBEZIUM_BASE_URL" env-default:"http://localhost:8080"`
}

type GRPCServerConfig struct {
	Port string `env:"GRPC_PORT"         env-default:"50123"`
}

func ParseConfig(path string) (*Config, error) {
	cfg := &Config{}
	postgresCfg := &postgrespool.Config{}
	grpcCfg := &GRPCServerConfig{}
	httpCfg := &HTTPServerConfig{}

	if err := cleanenv.ReadConfig(path, cfg); err != nil {
		return nil, err
	}

	if err := cleanenv.ReadConfig(path, postgresCfg); err != nil {
		return nil, fmt.Errorf("error config postgres: %w", err)
	}

	if err := cleanenv.ReadConfig(path, grpcCfg); err != nil {
		return nil, fmt.Errorf("error grpc postgres: %w", err)
	}

	if err := cleanenv.ReadConfig(path, httpCfg); err != nil {
		return nil, fmt.Errorf("error http postgres: %w", err)
	}

	cfg.PostgresConfig, cfg.GRPCServerConfig, cfg.HTTPServerConfig = postgresCfg, grpcCfg, httpCfg
	return cfg, nil
}
