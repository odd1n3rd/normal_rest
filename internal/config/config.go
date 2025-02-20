package config

import (
	"context"
	"fmt"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS" env-default:":12345"`
	DBAddr        string `env:"DB_ADDR"`
}

func Load() *Config {
	var cfg Config
	envconfig.MustProcess(context.Background(), &cfg)
	fmt.Println(cfg)
	return &cfg
}
