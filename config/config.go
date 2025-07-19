package config

import (
	"os"

	"github.com/spf13/viper"
)

var ConfigLoadedForTest bool

func Init() {
	if os.Getenv("ENVIRONMENT") == "test" {
		viper.SetConfigName("test")
	} else {
		viper.SetConfigName("application")
	}

	viper.SetConfigType("yml")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./../")
	viper.AddConfigPath("./../../")
	viper.AddConfigPath("./../../../")
	// For docker only
	viper.AddConfigPath("/app")

	err := viper.ReadInConfig()
	if err != nil {
		return
	}

	viper.AutomaticEnv()

	initAppConfig()
	initServerConfig()
	initDatabaseConfig()
	initLoggerConfig()
	initCacheConfig()
}

func InitForTest() {
	_ = os.Setenv("ENVIRONMENT", "test")
	if !ConfigLoadedForTest {
		Init()
	}
	ConfigLoadedForTest = true
}
