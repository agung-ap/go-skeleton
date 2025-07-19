package cache

import (
	"context"
	"fmt"
	"go-skeleton/config"
	"go-skeleton/pkg/logger"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

var (
	RedisClient *redis.Client
)

func Init(cfg config.CacheConfig) {
	// Create Redis client options
	options := &redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		PoolSize:     cfg.PoolSize,
		DialTimeout:  cfg.DialTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	// Create Redis client
	client := redis.NewClient(options)

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := client.Ping(ctx).Result()
	cancel()
	if err != nil {
		logger.Fatal("failed to connect to Redis", zap.Error(err))
	}

	RedisClient = client
}

func CloseCache() {
	if RedisClient != nil {
		if err := RedisClient.Close(); err != nil {
			logger.Fatal("failed to close Redis connection", zap.Error(err))
		}
	}
}

// GetClient returns the Redis client instance
func GetClient() *redis.Client {
	return RedisClient
}

// Set stores a key-value pair with optional expiration
func Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	return RedisClient.Set(ctx, key, value, expiration).Err()
}

// Get retrieves a value by key
func Get(ctx context.Context, key string) (string, error) {
	return RedisClient.Get(ctx, key).Result()
}

// Del deletes one or more keys
func Del(ctx context.Context, keys ...string) error {
	return RedisClient.Del(ctx, keys...).Err()
}

// Exists checks if a key exists
func Exists(ctx context.Context, keys ...string) (int64, error) {
	return RedisClient.Exists(ctx, keys...).Result()
}

// Expire sets a timeout on a key
func Expire(ctx context.Context, key string, expiration time.Duration) error {
	return RedisClient.Expire(ctx, key, expiration).Err()
}
