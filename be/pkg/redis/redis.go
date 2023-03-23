package redis

import (
	"github.com/just-hms/large-scale-multistructure-db/be/pkg/osext"

	redisdriver "github.com/go-redis/redis"
)

type Redis struct {
	Client *redisdriver.Client
}

func New() (*Redis, error) {
	redisAddr, err := osext.GetStringEnv("REDIS_ADDRESS")
	if err != nil {
		return nil, err
	}

	// it's ok if the password is empty
	redisPassword, _ := osext.GetStringEnv("REDIS_PASSWORD")

	return &Redis{
		Client: redisdriver.NewClient(&redisdriver.Options{
			Addr:     redisAddr,
			Password: redisPassword,
			DB:       0,
		}),
	}, nil
}
