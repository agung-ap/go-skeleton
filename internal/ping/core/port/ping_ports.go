package port

import "github.com/jmoiron/sqlx"

type PingRepository struct {
	DB *sqlx.DB
}
