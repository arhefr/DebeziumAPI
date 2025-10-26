package config

import (
	"fmt"
	"time"

	"debez/pkg/postgrespool"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	PostgresConfig   *postgrespool.Config
	GRPCServerConfig *GRPCServerConfig
	HTTPServerConfig *HTTPServerConfig
}

type HTTPServerConfig struct {
	Environment     string        `env:"ENV" env-default:"development"`
	Port            string        `env:"HTTP_PORT"         env-default:"8080"`
	Timeout         time.Duration `env:"HTTP_TIMEOUT" env-default:"30s"`
	DebeziumBaseURL string        `env:"DEBEZIUM_BASE_URL" env-default:"http://localhost:8080"`
}

type GRPCServerConfig struct {
	Port string `env:"GRPC_PORT"         env-default:"50123"`
}

func ParseConfig(path string) (*Config, error) {
	grpcCfg := &GRPCServerConfig{}
	httpCfg := &HTTPServerConfig{}
	postgresCfg := &postgrespool.Config{}

	if err := cleanenv.ReadConfig(path, httpCfg); err != nil {
		return nil, fmt.Errorf("ParseConfig: failed to parse config for HTTP server: %w", err)
	}

	if err := cleanenv.ReadConfig(path, grpcCfg); err != nil {
		return nil, fmt.Errorf("ParseConfig: failed to parse config for gRPC server: %w", err)
	}

	if err := cleanenv.ReadConfig(path, postgresCfg); err != nil {
		return nil, fmt.Errorf("ParseConfig: failed to parse config for postgres pool: %w", err)
	}

	cfg := &Config{
		PostgresConfig:   postgresCfg,
		GRPCServerConfig: grpcCfg,
		HTTPServerConfig: httpCfg,
	}
	return cfg, nil
}
