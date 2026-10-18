package config

import (
	"log/slog"

	"github.com/caarlos0/env/v11"
)

type Log struct {
	Level       slog.Level `env:"LOG_LEVEL" envDefault:"info"`
	ServiceName string     `env:"SERVICE_NAME" envDefault:"app"`
	Environment string     `env:"DEPLOYMENT_ENVIRONMENT" envDefault:"local"`
}

type Serve struct {
	Port      int `env:"PORT" envDefault:"8080"`
	DB        DBConfig
	Log       Log
	AppConfig AppConfig
}

type DBConfig struct {
	URL string `env:"DATABASE_URL,required"`
}

type AppConfig struct {
	AppMessage string `env:"APP_MESSAGE"`
	AppToken   string `env:"APP_TOKEN"`
}

func Load[T any]() (T, error) {
	return env.ParseAs[T]()
}
