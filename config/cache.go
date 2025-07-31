package config

import (
	"time"

	"github.com/spf13/viper"
)

type CacheConfig struct {
	Host         string
	Username     string
	Password     string
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
	DB           int
	Port         int
	PoolSize     int
}

var RedisCache CacheConfig

func initCacheConfig() {
	RedisCache = CacheConfig{
		Host:         viper.GetString("REDIS_HOST"),
		Username:     viper.GetString("REDIS_USERNAME"),
		Password:     viper.GetString("REDIS_PASSWORD"),
		DB:           viper.GetInt("REDIS_DB"),
		Port:         viper.GetInt("REDIS_PORT"),
		PoolSize:     viper.GetInt("REDIS_POOL_SIZE"),
		DialTimeout:  time.Duration(viper.GetDuration("REDIS_DIAL_TIMEOUT").Milliseconds()),
		ReadTimeout:  time.Duration(viper.GetDuration("REDIS_READ_TIMEOUT").Milliseconds()),
		WriteTimeout: time.Duration(viper.GetDuration("REDIS_WRITE_TIMEOUT").Milliseconds()),
		IdleTimeout:  time.Duration(viper.GetDuration("REDIS_IDLE_TIMEOUT").Milliseconds()),
	}
}
