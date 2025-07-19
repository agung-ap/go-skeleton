package database

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // file source driver
)

// MigrationManager handles database migrations
type MigrationManager struct {
	Directory string
}

// NewMigrationManager creates a new migration manager
func NewMigrationManager(directory string) *MigrationManager {
	return &MigrationManager{
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

	// Create up migration file
	upFileName := fmt.Sprintf("%s_%s.up.sql", timestamp, name)
	upFilePath := filepath.Join(mm.Directory, upFileName)

	if err := os.WriteFile(upFilePath, []byte("-- Write your UP migration SQL here\n"), 0600); err != nil {
		return fmt.Errorf("failed to create up migration file: %w", err)
	}

	// Create down migration file
	downFileName := fmt.Sprintf("%s_%s.down.sql", timestamp, name)
	downFilePath := filepath.Join(mm.Directory, downFileName)

	if err := os.WriteFile(downFilePath, []byte("-- Write your DOWN migration SQL here\n"), 0600); err != nil {
		return fmt.Errorf("failed to create down migration file: %w", err)
	}

	fmt.Printf("Created migration files:\n  %s\n  %s\n", upFilePath, downFilePath)
	return nil
}

// getMigrate creates a new migrate instance
func (mm *MigrationManager) getMigrate() (*migrate.Migrate, error) {
	// Create a new postgres driver
	driver, err := postgres.WithInstance(MigrationDB.DB, &postgres.Config{
		MigrationsTable: "schema_migrations",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create postgres driver: %w", err)
	}

	// Create a new migrate instance
	sourceURL := fmt.Sprintf("file://%s", mm.Directory)
	m, err := migrate.NewWithDatabaseInstance(sourceURL, "postgres", driver)
	if err != nil {
		return nil, fmt.Errorf("failed to create migrate instance: %w", err)
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
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	if sourceErr, dbErr := m.Close(); sourceErr != nil || dbErr != nil {
		fmt.Printf("Warning: failed to close migration instance: source=%v, db=%v\n", sourceErr, dbErr)
	}

	fmt.Println("Migrations applied successfully")
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
		return fmt.Errorf("failed to apply %d migrations: %w", steps, err)
	}

	fmt.Printf("Applied %d migrations successfully\n", steps)
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
		return fmt.Errorf("failed to roll back %d migrations: %w", steps, err)
	}

	fmt.Printf("Rolled back %d migrations successfully\n", steps)
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
		return fmt.Errorf("failed to roll back all migrations: %w", err)
	}

	fmt.Println("All migrations rolled back successfully")
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
		return fmt.Errorf("failed to migrate to version %d: %w", version, err)
	}

	fmt.Printf("Migrated to version %d successfully\n", version)
	return nil
}

// GetCurrentVersion returns the current migration version
func (mm *MigrationManager) GetCurrentVersion() (uint, bool, error) {
	m, err := mm.getMigrate()
	if err != nil {
		return 0, false, err
	}
	// Don't close migrate instance to avoid affecting main Migration connection
	// defer m.Close()

	return m.Version()
}
