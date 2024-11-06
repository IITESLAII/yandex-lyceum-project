package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"yandexlms/pkg/db/cache"
	"yandexlms/pkg/db/postgres"
)

type Config struct {
	postgres.Config
	cache.RedisConfig

	GRPCServerPort int `env:"GRPC_SERVER_PORT" env-default:"9090"`
	RestServerPort int `env:"REST_SERVER_PORT" env-default:"8080"`
}

func New() *Config {
	cfg := Config{}

	err := cleanenv.ReadConfig("./configs/local.env", &cfg)
	if err != nil {
		return nil
	}
	return &cfg
}
