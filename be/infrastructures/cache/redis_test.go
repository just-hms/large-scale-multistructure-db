package cache_test

import (
	"context"
	"testing"
	"time"

	"large-scale-multistructure-db/be/infrastructures/cache"

	"github.com/go-redis/redis/v8"
)

func TestRedisSetup(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := redis.NewClient(&redis.Options{
		Addr:     cache.RedisAddr,
		Password: cache.RedisPassword,
		DB:       0,
	})

	_, err := client.Ping(ctx).Result()

	if err != nil {
		t.Errorf("Failed to connect to Redis: %v", err)
	}
}
