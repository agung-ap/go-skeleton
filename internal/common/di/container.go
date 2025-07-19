package common

import (
	"go-skeleton/pkg/database"

	"github.com/jmoiron/sqlx"
)

// Container holds all the dependencies for the application
type Container struct {
	DB *sqlx.DB
}

// NewContainer creates a new dependency injection container
func NewContainer() Container {
	return Container{
		DB: database.DBConn,
	}
}
