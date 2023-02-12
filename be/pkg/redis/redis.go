package redis

import (
	redisdriver "github.com/go-redis/redis"
)

// TODO: put this in
const (
	DEFAULT_ADDRESS = "cache:6379"
	DEFAULT_PASWORD = ""
	DEFAULT_DB      = 0
)

type Redis struct {
	Client *redisdriver.Client
}

type RedisOptions struct {
	Address  string
	Password string
	DB       int
}

// get url and options as param
// add const

// get url and options as param

func New(opt *RedisOptions) *Redis {

	if opt.Address == "" {
		opt.Address = DEFAULT_ADDRESS
	}
	if opt.DB == 0 {
		opt.DB = DEFAULT_DB
	}
	if opt.Password == "" {
		opt.Password = DEFAULT_PASWORD
	}

	return &Redis{
		Client: redisdriver.NewClient(&redisdriver.Options{
			Addr:     opt.Address,
			Password: opt.Password,
			DB:       opt.DB,
		}),
	}
}
