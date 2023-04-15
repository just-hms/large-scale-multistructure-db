package redis

import (
	"fmt"

	"github.com/just-hms/large-scale-multistructure-db/be/pkg/env"

	redisdriver "github.com/go-redis/redis"
)

type Redis struct {
	Client *redisdriver.Client
}

func New() (*Redis, error) {

	redisHost, err := env.GetString("REDIS_HOST")
	if err != nil {
		redisHost = "localhost"
	}

	redisPort, err := env.GetInteger("REDIS_PORT")
	if err != nil {
		return nil, err
	}

	redisAddr := fmt.Sprintf("%s:%d", redisHost, redisPort)

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

func (r *Redis) Clear() error {
	_, err := r.Client.FlushDB().Result()
	return err
}
