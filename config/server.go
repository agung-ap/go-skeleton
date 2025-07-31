package config

import (
	"time"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var Server ServerConfig

func initServerConfig() {
	Server = ServerConfig{
		Port:         viper.GetInt("SERVER_PORT"),
		ReadTimeout:  time.Duration(viper.GetDuration("READ_TIMEOUT_MS").Milliseconds()),
		WriteTimeout: time.Duration(viper.GetDuration("WRITE_TIMEOUT_MS").Milliseconds()),
	}
}
