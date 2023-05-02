package redis

import (
	"fmt"

	redisdriver "github.com/go-redis/redis"
)

type Redis struct {
	Client *redisdriver.Client
}

func New(host string, port int, password string) (*Redis, error) {

	redisAddr := fmt.Sprintf("%s:%d", host, port)

	return &Redis{
		Client: redisdriver.NewClient(&redisdriver.Options{
			Addr:     redisAddr,
			Password: password,
			DB:       0,
		}),
	}, nil
}

func (r *Redis) Clear() error {
	_, err := r.Client.FlushAllAsync().Result()
	return err
}
