package redis

import (
	"github.com/just-hms/large-scale-multistructure-db/be/pkg/env"

	redisdriver "github.com/go-redis/redis"
)

type Redis struct {
	Client *redisdriver.Client
}

func New() (*Redis, error) {
	redisAddr, err := env.GetString("REDIS_ADDRESS")
	if err != nil {
		return nil, err
	}

	// it's ok if the password is empty
	redisPassword, _ := env.GetString("REDIS_PASSWORD")

	return &Redis{
		Client: redisdriver.NewClient(&redisdriver.Options{
			Addr:     redisAddr,
			Password: redisPassword,
			DB:       0,
		}),
	}, nil
}
