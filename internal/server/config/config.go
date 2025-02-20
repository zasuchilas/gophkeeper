package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	App        string `yaml:"app" env-default:"gophkeeper"`
	Env        string `yaml:"env" env-default:"local"`
	GRPCServer `yaml:"grpc_server"`
	PostgreSQL `yaml:"postgresql"`
	JWT        `yaml:"jwt"`
}

type GRPCServer struct {
	Address string `yaml:"address" env-default:"localhost:9999"`
}

type PostgreSQL struct {
	DSN string `yaml:"dsn" env-required:"true"`
}

type JWT struct {
	Secrets    []string      `yaml:"secrets" env-required:"true"`
	SessionTTL time.Duration `yaml:"session_ttl" env-default:"24h"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
