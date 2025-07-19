package config

import "github.com/spf13/viper"

type LoggerConfig struct {
	Level             string   `mapstructure:"LOG_LEVEL"`
	Development       bool     `mapstructure:"LOG_DEVELOPMENT"`
	DisableCaller     bool     `mapstructure:"LOG_DISABLE_CALLER"`
	DisableStacktrace bool     `mapstructure:"LOG_DISABLE_STACKTRACE"`
	Encoding          string   `mapstructure:"LOG_ENCODING"`
	OutputPaths       []string `mapstructure:"LOG_OUTPUT_PATHS"`
	ErrorOutputPaths  []string `mapstructure:"LOG_ERROR_OUTPUT_PATHS"`
}

var Logger LoggerConfig

func initLoggerConfig() {
	Logger.Level = getStringOrDefault("LOG_LEVEL", "info")
	Logger.Development = getBoolOrDefault("LOG_DEVELOPMENT", false)
	Logger.DisableCaller = getBoolOrDefault("LOG_DISABLE_CALLER", false)
	Logger.DisableStacktrace = getBoolOrDefault("LOG_DISABLE_STACKTRACE", false)
	Logger.Encoding = getStringOrDefault("LOG_ENCODING", "json")
	Logger.OutputPaths = getStringSliceOrDefault("LOG_OUTPUT_PATHS", []string{"stdout"})
	Logger.ErrorOutputPaths = getStringSliceOrDefault("LOG_ERROR_OUTPUT_PATHS", []string{"stderr"})
}

func getStringOrDefault(key, defaultValue string) string {
	if value := viper.GetString(key); value != "" {
		return value
	}
	return defaultValue
}

func getBoolOrDefault(key string, defaultValue bool) bool {
	if value := viper.GetString(key); value != "" {
		return value == "true"
	}
	return defaultValue
}

func getStringSliceOrDefault(key string, defaultValue []string) []string {
	if value := viper.GetString(key); value != "" {
		return []string{value}
	}
	return defaultValue
}
