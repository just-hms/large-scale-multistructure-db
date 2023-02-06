package redis

import (
	redisdriver "github.com/go-redis/redis"
)

// TODO: put this in
const RedisAddr = "cache:6379"
const RedisPassword = ""

type Redis struct {
	Client *redisdriver.Client
}

// get url and options as param
// add const

// get url and options as param

func New() *Redis {

	return &Redis{
		Client: redisdriver.NewClient(&redisdriver.Options{
			Addr:     RedisAddr,
			Password: RedisPassword,
			DB:       0,
		}),
	}
}
