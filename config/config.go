package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

var ConfigLoadedForTest bool

func Init() {
	// Determine the current environment. Default to "development" when not set.
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "development"
		_ = os.Setenv("ENVIRONMENT", env)
	}

	// Load the environment-specific .env file (e.g. env/development.env).
	var b strings.Builder
	b.WriteString("env/")
	b.WriteString(env)
	b.WriteString(".env")

	_ = gotenv.Load(b.String())

	// Bind environment variables to Viper
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
