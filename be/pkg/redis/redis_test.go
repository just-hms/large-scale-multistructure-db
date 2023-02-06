package redis_test

import (
	"large-scale-multistructure-db/be/pkg/redis"
	"testing"
)

func TestRedisSetup(t *testing.T) {

	redis := redis.New()

	_, err := redis.Client.Ping().Result()

	if err != nil {
		t.Errorf("Failed to connect to Redis: %v", err)
	}
}
