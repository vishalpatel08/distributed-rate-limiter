package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/vishalpatel08/distributed-rate-limiter/internal/config"
)

type Repository struct {
	cfg    *config.Config
	client *redis.Client
}

func NewRepository(cfg *config.Config, client *redis.Client) *Repository {
	return &Repository{
		cfg:    cfg,
		client: client,
	}
}

func (r Repository) ConsumeToken(clientID string) (bool, int, error) {

	keys := []string{
		fmt.Sprintf("bucket:%s:tokens", clientID),
		fmt.Sprintf("bucket:%s:last_refill", clientID),
	}

	args := []interface{}{
		r.cfg.BucketCapacity,
		r.cfg.RefillRate,
		time.Now().Unix(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	result, err := r.client.Eval(
		ctx,
		TokenBucketLua,
		keys,
		args...,
	).Result()

	if err != nil {
		return false, 0, err
	}

	values := result.([]interface{})

	allowed := values[0].(int64) == 1
	remaining := int(values[1].(int64))

	return allowed, remaining, nil
}
