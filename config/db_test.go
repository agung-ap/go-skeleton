package config

import (
	"os"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func setupDatabaseConfigForTest() {
	viper.Reset()
	viper.Set("DB_DRIVER", "postgres")
	viper.Set("DB_NAME", "test_db")
	viper.Set("DB_HOST", "localhost")
	viper.Set("DB_USER", "testuser")
	viper.Set("DB_PASSWORD", "testpass")
	viper.Set("DB_PORT", "5432")
	viper.Set("DB_POOL_SIZE", "10")
	viper.Set("DB_READ_TIMEOUT_MS", "1000")
	viper.Set("DB_WRITE_TIMEOUT_MS", "2000")
	viper.Set("DB_CONNECTION_MAX_LIFETIME_MINUTE", "30")
}

func TestInitDatabaseConfig(t *testing.T) {
	setupDatabaseConfigForTest()

	initDatabaseConfig()

	assert.Equal(t, "postgres", Database.DriverName)
	assert.Equal(t, "test_db", Database.Name)
	assert.Equal(t, "localhost", Database.Host)
	assert.Equal(t, "testuser", Database.User)
	assert.Equal(t, "testpass", Database.Password)
	assert.Equal(t, 5432, Database.Port)
	assert.Equal(t, 10, Database.MaxPoolSize)
	assert.Equal(t, time.Duration(1000)*time.Millisecond, Database.ReadTimeout)
	assert.Equal(t, time.Duration(2000)*time.Millisecond, Database.WriteTimeout)
	assert.Equal(t, time.Duration(30)*time.Minute, Database.ConnectionMaxLifeTime)
}

func TestInitDatabaseConfig_WithTestEnvironment(t *testing.T) {
	// Skip this test if not in test environment
	if os.Getenv("ENVIRONMENT") != "test" {
		t.Skip("Skipping test environment test")
	}

	// Reset viper and load test configuration
	viper.Reset()
	t.Setenv("ENVIRONMENT", "test")
	Init()

	assert.Equal(t, "postgres", Database.DriverName)
	assert.Equal(t, "test-db", Database.Name)
	assert.Equal(t, "postgres-db-test", Database.Host)
	assert.Equal(t, "postgres", Database.User)
	assert.Equal(t, "postgres", Database.Password)
	assert.Equal(t, 5432, Database.Port)
	assert.Equal(t, 20, Database.MaxPoolSize)
	assert.Equal(t, time.Duration(200)*time.Millisecond, Database.ReadTimeout)
	assert.Equal(t, time.Duration(200)*time.Millisecond, Database.WriteTimeout)
	assert.Equal(t, time.Duration(20)*time.Minute, Database.ConnectionMaxLifeTime)
}

func TestDatabaseConfig_ConnectionURL(t *testing.T) {
	config := DatabaseConfig{
		DriverName: "postgres",
		User:       "testuser",
		Password:   "testpass",
		Host:       "localhost",
		Port:       5432,
		Name:       "testdb",
	}

	expected := "postgres://testuser:testpass@localhost:5432/testdb?sslmode=disable"
	result := config.ConnectionURL()

	assert.Equal(t, expected, result)
}

func TestDatabaseConfig_ConnectionURL_WithSpecialCharacters(t *testing.T) {
	config := DatabaseConfig{
		DriverName: "postgres",
		User:       "testuser",
		Password:   "test@pass#word",
		Host:       "localhost",
		Port:       5432,
		Name:       "testdb",
	}

	result := config.ConnectionURL()

	// Should URL encode the password
	assert.Contains(t, result, "test%40pass%23word")
	assert.Contains(t, result, "postgres://testuser:")
	assert.Contains(t, result, "@localhost:5432/testdb?sslmode=disable")
}
