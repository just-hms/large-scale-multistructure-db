package redis_test

import (
	"testing"

	"github.com/just-hms/large-scale-multistructure-db/be/pkg/redis"

	"github.com/stretchr/testify/assert"
)

func TestRedisSetup(t *testing.T) {

	r, err := redis.New()
	assert.Nil(t, err)

	_, err = r.Client.Ping().Result()

	assert.Nil(t, err)
}
