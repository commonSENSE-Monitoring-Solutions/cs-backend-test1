package config

import (
	"fmt"
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Database struct {
	Host     string `env:"POSTGRES_HOST"`
	Port     int    `env:"POSTGRES_PORT"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
}

type Logger struct {
	LogLevel int8 `env:"LOG_LEVEL"`
}

type Config struct {
	Db     Database
	Logger Logger
}

func LoadAndParse() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("error parsing .env variables to struct: %w", err)
	}

	return &cfg, nil
}
