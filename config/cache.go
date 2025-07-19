package config

import (
	"time"
)

type CacheConfig struct {
	Host         string
	Port         int
	PoolSize     int
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

var RedisCache CacheConfig

func initCacheConfig() {
	RedisCache = CacheConfig{
		Host:         mustGetString("REDIS_HOST"),
		Port:         mustGetInt("REDIS_PORT"),
		PoolSize:     mustGetInt("REDIS_POOL_SIZE"),
		DialTimeout:  mustGetDurationMs("REDIS_DIAL_TIMEOUT"),
		ReadTimeout:  mustGetDurationMs("REDIS_READ_TIMEOUT"),
		WriteTimeout: mustGetDurationMs("REDIS_WRITE_TIMEOUT"),
		IdleTimeout:  mustGetDurationMs("REDIS_IDLE_TIMEOUT"),
	}
}
