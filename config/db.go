package config

import (
	"fmt"
	"net/url"
	"time"
)

type DatabaseConfig struct {
	DriverName            string
	Host                  string
	User                  string
	Password              string
	Port                  int
	MaxPoolSize           int
	ReadTimeout           time.Duration
	WriteTimeout          time.Duration
	ConnectionMaxOpen     int
	ConnectionMaxIdle     int
	ConnectionMaxLifeTime time.Duration
}

var Database DatabaseConfig

func initDatabaseConfig() {
	Database = DatabaseConfig{
		DriverName:            mustGetString("DB_DRIVER"),
		Host:                  mustGetString("DB_HOST"),
		User:                  mustGetString("DB_USER"),
		Password:              mustGetString("DB_PASSWORD"),
		Port:                  mustGetInt("DB_PORT"),
		MaxPoolSize:           mustGetInt("DB_POOL_SIZE"),
		ReadTimeout:           mustGetDurationMs("DB_READ_TIMEOUT_MS"),
		WriteTimeout:          mustGetDurationMs("DB_WRITE_TIMEOUT_MS"),
		ConnectionMaxLifeTime: mustGetDurationMinute("DB_CONNECTION_MAX_LIFETIME_MINUTE"),
	}
}

func (dc DatabaseConfig) ConnectionURL() string {
	return fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable",
		dc.DriverName,
		dc.User,
		url.QueryEscape(dc.Password),
		dc.Host,
		dc.Port,
		dc.DriverName,
	)
}
