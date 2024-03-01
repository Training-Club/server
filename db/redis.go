package db

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"tc-server/config"
	"time"
)

// GetRedisContext returns a pre-configured context
// used specifically for Redis cache processes.
func GetRedisContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 3*time.Second)
}

// InitRedis establishes a new connection to the Redis cache.
func InitRedis(conf *config.CacheConfig) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{Addr: conf.Address, Password: conf.Password, DB: conf.DBID})
	pong := rdb.Ping(context.Background())

	if pong.Err() != nil {
		return nil, pong.Err()
	}

	return rdb, nil
}

func SetCacheValue[K any](
	params RedisParams,
	key string,
	value K,
	ttl int,
) (string, error) {
	if params.RedisClient == nil {
		return "", fmt.Errorf("redis client is nil")
	}

	ctx, cancel := GetRedisContext()
	defer cancel()

	result := params.RedisClient.Set(ctx, key, value, time.Duration(ttl)*time.Second)
	if result.Err() != nil {
		return "", result.Err()
	}

	return result.Result()
}

func GetCacheValue(params RedisParams, key string) (string, error) {
	if params.RedisClient == nil {
		return "", fmt.Errorf("redis client is nil")
	}

	ctx, cancel := GetRedisContext()
	defer cancel()

	result := params.RedisClient.Get(ctx, key)
	if result.Err() != nil {
		return "", result.Err()
	}

	return result.Result()
}

func DeleteCacheValue(params RedisParams, key string) (int64, error) {
	if params.RedisClient == nil {
		return -1, fmt.Errorf("redis client is nil")
	}

	ctx, cancel := GetRedisContext()
	defer cancel()

	result := params.RedisClient.Del(ctx, key)
	if result.Err() != nil {
		return -1, result.Err()
	}

	return result.Result()
}
