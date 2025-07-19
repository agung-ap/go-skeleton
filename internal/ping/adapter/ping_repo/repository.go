package pingrepo

import (
	"github.com/jmoiron/sqlx"
)

type PingRepository struct {
	db *sqlx.DB
}

func CreateEnhancedRepository(db *sqlx.DB) PingRepository {
	return PingRepository{
		db: db,
	}
}

func (r PingRepository) Ping() string {
	if r.db != nil {
		return "pong from enhanced database repository"
	}

	return "pong from enhanced repository"
}
