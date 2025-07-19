package pingrepo

import (
	"go-skeleton/pkg/logger"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type PingRepository struct {
	db    *sqlx.DB
	cache *redis.Client
}

func CreateEnhancedRepository(db *sqlx.DB, cache *redis.Client) PingRepository {
	return PingRepository{
		db:    db,
		cache: cache,
	}
}

func (r PingRepository) Ping() string {
	if r.db != nil {
		logger.Info("Pinging database repository")
	}

	if r.cache != nil {
		logger.Info("Pinging cache repository")
	}

	return "pong from repository"
}
