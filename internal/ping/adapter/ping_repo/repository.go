package pingrepo

import (
	"go-skeleton/internal/ping/core/domain"
	"go-skeleton/pkg/logger"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type PingRepository struct {
	db    *sqlx.DB
	cache *redis.Client
}

func NewPingRepository(db *sqlx.DB, cache *redis.Client) PingRepository {
	return PingRepository{
		db:    db,
		cache: cache,
	}
}

func (r PingRepository) Ping() domain.Ping {
	if r.db != nil {
		logger.Info("Pinging database repository")
	}

	if r.cache != nil {
		logger.Info("Pinging cache repository")
	}

	result := PingResponse{
		Message: "ping from repository",
	}

	return result.ToDomain()
}
