package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddress string
	DBAddr        string
}

func Load() *Config {
	var cfg Config
	// envconfig.MustProcess(context.Background(), &cfg)

	if err := godotenv.Load(); err != nil {
		log.Fatal(map[string]string{"environment error": err.Error()})
	}
	cfg.ServerAddress = os.Getenv("SERVER_ADDRESS")
	cfg.DBAddr = os.Getenv("DB_DSN")
	// fmt.Println(cfg)
	return &cfg
}
