package database

import (
	"fmt"
	"go-skeleton/config"
	"go-skeleton/pkg/logger"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
	"go.uber.org/zap"
)

var (
	DBConn      *sqlx.DB
	MigrationDB *sqlx.DB
)

func Init(cfg config.DatabaseConfig) {
	// Build connection string for PostgreSQL
	dsn := cfg.ConnectionURL()

	// SQL connection for all operations
	db, err := sqlx.Connect(cfg.DriverName, dsn)
	if err != nil {
		logger.Fatal("failed to connect to database with sqlx", zap.Error(err))
	}

	// DB pool configuration
	db.SetMaxOpenConns(cfg.ConnectionMaxOpen)
	db.SetMaxIdleConns(cfg.ConnectionMaxIdle)
	db.SetConnMaxLifetime(cfg.ConnectionMaxLifeTime)

	DBConn = db
}

func CloseDB() {
	if DBConn.DB != nil {
		if err := DBConn.DB.Close(); err != nil {
			logger.Fatal("failed to close database connection", zap.Error(err))
		}
	}
}

func InitMigration(cfg config.DatabaseConfig) (*MigrationManager, error) {
	// Create a separate database connection specifically for migrations
	// Build connection string for PostgreSQL
	dsn := cfg.ConnectionURL()

	// SQL connection for all operations
	db, err := sqlx.Connect(cfg.DriverName, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database with sqlx: %w", err)
	}

	// DB pool configuration
	db.SetMaxOpenConns(cfg.ConnectionMaxOpen)
	db.SetMaxIdleConns(cfg.ConnectionMaxIdle)
	db.SetConnMaxLifetime(cfg.ConnectionMaxLifeTime)

	MigrationDB = db

	// Ensure migration connection is closed after migrations
	defer func() {
		if MigrationDB.DB != nil {
			_ = MigrationDB.DB.Close()
		}
	}()

	// Create a migration manager with the separate connection
	migrationsDir := "migrations"
	migrationManager := NewMigrationManager(migrationsDir)

	// Apply all migrations using the separate connection
	return migrationManager, nil
}
