package redis_test

import (
	"testing"

	"github.com/just-hms/large-scale-multistructure-db/be/config"
	"github.com/just-hms/large-scale-multistructure-db/be/pkg/redis"

	"github.com/stretchr/testify/assert"
)

func TestRedisSetup(t *testing.T) {

	cfg, err := config.NewConfig()
	assert.Nil(t, err)

	r, err := redis.New(cfg.Redis.Host, cfg.Redis.Port, cfg.Redis.Password)
	assert.Nil(t, err)

	_, err = r.Client.Ping().Result()

	assert.Nil(t, err)
}
