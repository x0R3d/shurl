package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env 		string `yaml:"env" env-default:"local"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer 	`yaml:"http_server"`
}

type HTTPServer struct {
	Address string `yaml:"address" env-default:"0.0.0.0:8080"`
	Timeout time.Duration `yaml:"timeout" env-default:"5s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("config file not found: %v", err)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	return &cfg
}