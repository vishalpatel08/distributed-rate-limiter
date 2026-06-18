package storage

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/vishalpatel08/distributed-rate-limiter/internal/config"
)

var Ctx = context.Background()

func NewRadisClient(cfg *config.Config) *redis.Client {
	addr := fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort)

	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	err := client.Ping(Ctx).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to redis: %v", err))
	}

	return client
}
