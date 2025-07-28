package container

import (
	"go-skeleton/pkg/cache"
	"go-skeleton/pkg/database"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

// Container holds all the dependencies for the application
type Container struct {
	DB    *sqlx.DB
	Cache *redis.Client
}

// NewContainer creates a new dependency injection container
func NewContainer() Container {
	return Container{
		DB:    database.DBConn,
		Cache: cache.RedisClient,
	}
}
