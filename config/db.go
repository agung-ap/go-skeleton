package config

import (
	"fmt"
	"net/url"
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
		DriverName:            viper.GetString("DB_DRIVER"),
		Name:                  viper.GetString("DB_NAME"),
		Host:                  viper.GetString("DB_HOST"),
		User:                  viper.GetString("DB_USER"),
		Password:              viper.GetString("DB_PASSWORD"),
		Port:                  viper.GetInt("DB_PORT"),
		MaxPoolSize:           viper.GetInt("DB_POOL_SIZE"),
		ReadTimeout:           time.Duration(viper.GetInt("DB_READ_TIMEOUT_MS")) * time.Millisecond,
		WriteTimeout:          time.Duration(viper.GetInt("DB_WRITE_TIMEOUT_MS")) * time.Millisecond,
		ConnectionMaxLifeTime: time.Duration(viper.GetInt("DB_CONNECTION_MAX_LIFETIME_MINUTE")) * time.Minute,
	}
}

func (dc DatabaseConfig) ConnectionURL() string {
	return fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable",
		dc.DriverName,
		dc.User,
		url.QueryEscape(dc.Password),
		dc.Host,
		dc.Port,
		dc.Name,
	)
}
