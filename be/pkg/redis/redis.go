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

	redisHost := env.GetStringWithDefault("REDIS_HOST", "localhost")
	redisPassword := env.GetStringWithDefault("REDIS_PASSWORD", "")
	redisPort, err := env.GetInt("REDIS_PORT")
	if err != nil {
		return nil, err
	}

	redisAddr := fmt.Sprintf("%s:%d", redisHost, redisPort)

	return &Redis{
		Client: redisdriver.NewClient(&redisdriver.Options{
			Addr:     redisAddr,
			Password: redisPassword,
			DB:       0,
		}),
	}, nil
}

func (r *Redis) Clear() error {
	_, err := r.Client.FlushAllAsync().Result()
	return err
}
