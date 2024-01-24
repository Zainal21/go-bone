package caching

import (
	"context"
	"fmt"

	"github.com/Zainal21/go-bone/pkg/config"
	"github.com/Zainal21/go-bone/pkg/logger"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func GetRedisClient(ctx context.Context, conf *config.Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.RedisHost,
		Password: conf.RedisPassword,
		DB:       conf.RedisDB,
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("unable to connect to Redis: %v", err)
	}
	logger.Info("Redis Connected")
	return rdb, nil
}

func SetRedisKey(ctx context.Context, key string, value interface{}) error {
	return rdb.Set(ctx, key, value, 0).Err()
}

func GetRedisKey(ctx context.Context, key string) (string, error) {
	return rdb.Get(ctx, key).Result()
}
