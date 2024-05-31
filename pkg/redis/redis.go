package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"golang-crud-2024/config"
)

type RedisClient struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisClient(cfg config.Config) (*RedisClient, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", cfg.MySqlHost, cfg.RedisPort),
	})
	ctx := context.Background()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &RedisClient{client: rdb, ctx: ctx}, nil
}

func (r *RedisClient) Set(key string, value interface{}) error {
	return r.client.Set(r.ctx, key, value, 0).Err()
}

func (r *RedisClient) Get(key string) (string, error) {
	return r.client.Get(r.ctx, key).Result()
}
