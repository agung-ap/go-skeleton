package database

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	x "go-skeleton/pkg/errors/entity"
	"go-skeleton/pkg/logger"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // file source driver
	"go.uber.org/zap"
)

// MigrationManager handles database migrations
type MigrationManager struct {
	Directory string
}

// NewMigrationManager creates a new migration manager
func NewMigrationManager(directory string) MigrationManager {
	return MigrationManager{
		Directory: directory,
	}
}

// CreateMigration creates a new migration file with the given name
func (mm *MigrationManager) CreateMigration(name string) error {
	// Ensure migrations directory exists
	if err := os.MkdirAll(mm.Directory, 0755); err != nil {
		return fmt.Errorf("failed to create migrations directory: %w", err)
	}

	// Generate timestamp for migration version
	timestamp := time.Now().Format("20060102150405")

	// Create up migration file using strings.Builder to minimize allocations
	var upNameBuilder strings.Builder
	upNameBuilder.Grow(len(timestamp) + 1 + len(name) + len(".up.sql"))
	upNameBuilder.WriteString(timestamp)
	upNameBuilder.WriteByte('_')
	upNameBuilder.WriteString(name)
	upNameBuilder.WriteString(".up.sql")
	upFileName := upNameBuilder.String()
	upFilePath := filepath.Join(mm.Directory, upFileName)

	if err := os.WriteFile(upFilePath, []byte("-- Write your UP migration SQL here\n"), 0600); err != nil {
		return fmt.Errorf("failed to create up migration file: %w", err)
	}

	// Create down migration file using strings.Builder
	var downNameBuilder strings.Builder
	downNameBuilder.Grow(len(timestamp) + 1 + len(name) + len(".down.sql"))
	downNameBuilder.WriteString(timestamp)
	downNameBuilder.WriteByte('_')
	downNameBuilder.WriteString(name)
	downNameBuilder.WriteString(".down.sql")
	downFileName := downNameBuilder.String()
	downFilePath := filepath.Join(mm.Directory, downFileName)

	if err := os.WriteFile(downFilePath, []byte("-- Write your DOWN migration SQL here\n"), 0600); err != nil {
		return x.Wrap(err, "failed to create down migration file")
	}

	logger.Info("Created migration files",
		zap.String("up_file", upFilePath),
		zap.String("down_file", downFilePath))
	return nil
}

// getMigrate creates a new migrate instance
func (mm *MigrationManager) getMigrate() (*migrate.Migrate, error) {
	pgConfig := postgres.Config{
		MigrationsTable: "schema_migrations",
	}

	// Create a new postgres driver
	driver, err := postgres.WithInstance(MigrationDB.DB, &pgConfig)
	if err != nil {
		return nil, x.Wrap(err, "failed to create postgres driver")
	}

	// Create a new migrate instance
	sourceURL := fmt.Sprintf("file://%s", mm.Directory)
	m, err := migrate.NewWithDatabaseInstance(sourceURL, "postgres", driver)
	if err != nil {
		return nil, x.Wrap(err, "failed to create migrate instance")
	}

	return m, nil
}

// ApplyMigrations applies all migrations
func (mm *MigrationManager) ApplyMigrations() error {
	m, err := mm.getMigrate()
	if err != nil {
		return err
	}

	// Apply all migrations
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return x.Wrap(err, "failed to apply migrations")
	}

	if sourceErr, dbErr := m.Close(); sourceErr != nil || dbErr != nil {
		logger.Warn("Warning: failed to close migration instance", zap.Error(sourceErr), zap.Error(dbErr))
	}

	// Close the migration database connection
	CloseMigrationDB()

	return nil
}

// ApplyMigrationsSteps applies a specific number of migrations
func (mm *MigrationManager) ApplyMigrationsSteps(steps int) error {
	m, err := mm.getMigrate()
	if err != nil {
		return err
	}
	defer m.Close()

	// Apply specific number of migrations
	if err := m.Steps(steps); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return x.Wrap(err, fmt.Sprintf("failed to apply %d migrations", steps))
	}

	// Close the migration database connection
	CloseMigrationDB()

	logger.Info("Applied migrations successfully", zap.Int("steps", steps))
	return nil
}

// RollbackMigration rolls back the last applied migration
func (mm *MigrationManager) RollbackMigration() error {
	return mm.RollbackMigrationsSteps(1)
}

// RollbackMigrationsSteps rolls back a specific number of migrations
func (mm *MigrationManager) RollbackMigrationsSteps(steps int) error {
	m, err := mm.getMigrate()
	if err != nil {
		return err
	}
	// Don't close migrate instance to avoid affecting main Migration connection
	// defer m.Close()

	// Roll back specific number of migrations
	if err := m.Steps(-steps); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return x.Wrap(err, fmt.Sprintf("failed to roll back %d migrations", steps))
	}

	// Close the migration database connection
	CloseMigrationDB()

	logger.Info("Rolled back migrations successfully", zap.Int("steps", steps))
	return nil
}

// RollbackAllMigrations rolls back all applied migrations
func (mm *MigrationManager) RollbackAllMigrations() error {
	m, err := mm.getMigrate()
	if err != nil {
		return err
	}
	// Don't close migrate instance to avoid affecting main Migration connection
	// defer m.Close()

	// Roll back all migrations
	if err := m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return x.Wrap(err, "failed to roll back all migrations")
	}

	// Close the migration database connection
	CloseMigrationDB()

	return nil
}

// MigrateTo migrates to a specific version
func (mm *MigrationManager) MigrateTo(version uint) error {
	m, err := mm.getMigrate()
	if err != nil {
		return err
	}
	// Don't close migrate instance to avoid affecting main Migration connection
	// defer m.Close()

	// Migrate to specific version
	if err := m.Migrate(version); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return x.Wrap(err, fmt.Sprintf("failed to migrate to version %d", version))
	}

	// Close the migration database connection
	CloseMigrationDB()

	logger.Info("Migrated to version successfully", zap.Uint("version", version))
	return nil
}

// GetCurrentVersion returns the current migration version
func (mm *MigrationManager) GetCurrentVersion() (uint, bool, error) {
	m, err := mm.getMigrate()
	if err != nil {
		return 0, false, x.Wrap(err, "failed to get current version")
	}

	return m.Version()
}
