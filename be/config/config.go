package config

import (
	"errors"
	"path/filepath"
	"runtime"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		JWT
		Mongo
		Redis
		Geocoding
	}

	JWT struct {
		TokenLifespan time.Duration `env-required:"true" env:"TOKEN_LIFE_SPAN"`
		ApiSecret     string        `env-required:"true" env:"TOKEN_API_SECRET"`
	}

	Mongo struct {
		Host string `env:"MONGO_HOST" env-default:"localhost"`
		Port int    `env-required:"true" env:"MONGO_PORT"`
	}

	Redis struct {
		Host     string `env:"REDIS_HOST" env-default:"localhost"`
		Port     int    `env-required:"true" env:"REDIS_PORT"`
		Password string `env:"REDIS_PASSWORD" env-default:""`
	}
	Geocoding struct {
		Apikey string `env-required:"true" env:"GEOCODE_API_SECRET"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	_, caller, _, ok := runtime.Caller(0)
	if !ok {
		return nil, errors.New("error retrevieng the .env file")
	}
	envPath := filepath.Join(filepath.Dir(caller), "../../.env")

	cfg := &Config{}
	err := cleanenv.ReadConfig(envPath, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
