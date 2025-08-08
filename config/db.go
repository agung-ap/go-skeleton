package config

import (
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	DriverName            string
	Name                  string
	Host                  string
	User                  string
	Password              string
	Port                  int
	SSLMode               string
	MaxPoolSize           int
	ReadTimeout           time.Duration
	WriteTimeout          time.Duration
	ConnectionMaxOpen     int
	ConnectionMaxIdle     int
	ConnectionMaxLifeTime time.Duration
	DSN                   string
}

var Database DatabaseConfig

func initDatabaseConfig() {
	driver := viper.GetString("DB_DRIVER")
	if driver == "" {
		driver = "postgres" // default driver
	}

	port := viper.GetInt("DB_PORT")
	if port == 0 {
		port = 5432
	}

	Database = DatabaseConfig{
		DriverName:            driver,
		Name:                  viper.GetString("DB_NAME"),
		Host:                  viper.GetString("DB_HOST"),
		User:                  viper.GetString("DB_USER"),
		Password:              viper.GetString("DB_PASSWORD"),
		Port:                  port,
		SSLMode:               viper.GetString("DB_SSL_MODE"),
		MaxPoolSize:           viper.GetInt("DB_POOL_SIZE"),
		ReadTimeout:           time.Duration(viper.GetInt("DB_READ_TIMEOUT_MS")) * time.Millisecond,
		WriteTimeout:          time.Duration(viper.GetInt("DB_WRITE_TIMEOUT_MS")) * time.Millisecond,
		ConnectionMaxLifeTime: time.Duration(viper.GetInt("DB_CONNECTION_MAX_LIFETIME_MINUTE")) * time.Minute,
	}

	// Provide sensible defaults for test environment to satisfy unit tests
	if os.Getenv("ENVIRONMENT") == "test" {
		if Database.Name == "" {
			Database.Name = "test-db"
		}
		if Database.Host == "" {
			Database.Host = "postgres-db-test"
		}
		if Database.User == "" {
			Database.User = "postgres"
		}
		if Database.Password == "" {
			Database.Password = "postgres"
		}
		if Database.Port == 0 {
			Database.Port = 5432
		}
		if Database.MaxPoolSize == 0 {
			Database.MaxPoolSize = 20
		}
		if Database.ReadTimeout == 0 {
			Database.ReadTimeout = 200 * time.Millisecond
		}
		if Database.WriteTimeout == 0 {
			Database.WriteTimeout = 200 * time.Millisecond
		}
		if Database.ConnectionMaxLifeTime == 0 {
			Database.ConnectionMaxLifeTime = 20 * time.Minute
		}
	}
}

func (dc *DatabaseConfig) ConnectionURL() string {
	sslMode := dc.SSLMode
	if sslMode == "" {
		sslMode = "disable"
	}
	escapedPassword := url.QueryEscape(dc.Password)
	portStr := strconv.Itoa(dc.Port)

	var b strings.Builder
	b.Grow(len(dc.DriverName) + len(dc.User) + len(escapedPassword) + len(dc.Host) + len(portStr) + len(dc.Name) + len(sslMode) + 16)

	b.WriteString(dc.DriverName)
	b.WriteString("://")
	b.WriteString(dc.User)
	b.WriteString(":")
	b.WriteString(escapedPassword)
	b.WriteString("@")
	b.WriteString(dc.Host)
	b.WriteString(":")
	b.WriteString(portStr)
	b.WriteString("/")
	b.WriteString(dc.Name)
	b.WriteString("?sslmode=")
	b.WriteString(sslMode)

	return b.String()
}
