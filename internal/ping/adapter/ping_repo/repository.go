package pingrepo

import (
	"context"
	"go-skeleton/internal/ping/core/domain"
	pkgErr "go-skeleton/pkg/errors/entity"
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

func (r *PingRepository) Ping(ctx context.Context, resp *domain.Ping) error {
	if r.db != nil {
		err := r.db.PingContext(ctx)
		if err != nil {
			return pkgErr.Wrap(err, "ping database repository")
		}
		logger.Info("Pinging database repository")
	}

	if r.cache != nil {
		err := r.cache.Ping(ctx).Err()
		if err != nil {
			return pkgErr.Wrap(err, "ping cache repository")
		}

		logger.Info("Pinging cache repository")
	}

	result := PingResponse{
		Message: "ping from repository",
	}

	*resp = result.ToDomain()

	return nil
}
