package config

import (
	"strings"

	"github.com/spf13/viper"
)

type LoggerConfig struct {
	OutputPaths       []string `mapstructure:"LOG_OUTPUT_PATHS"`
	ErrorOutputPaths  []string `mapstructure:"LOG_ERROR_OUTPUT_PATHS"`
	Level             string   `mapstructure:"LOG_LEVEL"`
	Encoding          string   `mapstructure:"LOG_ENCODING"`
	Development       bool     `mapstructure:"LOG_DEVELOPMENT"`
	DisableCaller     bool     `mapstructure:"LOG_DISABLE_CALLER"`
	DisableStacktrace bool     `mapstructure:"LOG_DISABLE_STACKTRACE"`
}

var Logger LoggerConfig

var (
	defaultOutputPaths      = []string{"stdout"}
	defaultErrorOutputPaths = []string{"stderr"}
)

func initLoggerConfig() {
	// LOG_LEVEL
	level := viper.GetString("LOG_LEVEL")
	if level == "" {
		level = "info"
	}
	Logger.Level = level

	// LOG_DEVELOPMENT
	devStr := viper.GetString("LOG_DEVELOPMENT")
	if devStr == "" {
		Logger.Development = false
	} else {
		Logger.Development = devStr == "true"
	}

	// LOG_DISABLE_CALLER
	disableCallerStr := viper.GetString("LOG_DISABLE_CALLER")
	if disableCallerStr == "" {
		Logger.DisableCaller = false
	} else {
		Logger.DisableCaller = disableCallerStr == "true"
	}

	// LOG_DISABLE_STACKTRACE
	disableStackStr := viper.GetString("LOG_DISABLE_STACKTRACE")
	if disableStackStr == "" {
		Logger.DisableStacktrace = false
	} else {
		Logger.DisableStacktrace = disableStackStr == "true"
	}

	// LOG_ENCODING
	encoding := viper.GetString("LOG_ENCODING")
	if encoding == "" {
		encoding = "json"
	}
	Logger.Encoding = encoding

	// LOG_OUTPUT_PATHS
	outputPaths := viper.GetStringSlice("LOG_OUTPUT_PATHS")
	if len(outputPaths) == 0 {
		if s := viper.GetString("LOG_OUTPUT_PATHS"); s != "" {
			parts := strings.Split(s, ",")
			for i := range parts {
				parts[i] = strings.TrimSpace(parts[i])
			}
			outputPaths = parts
		}
	}
	if len(outputPaths) == 0 {
		outputPaths = append([]string(nil), defaultOutputPaths...)
	}
	Logger.OutputPaths = outputPaths

	// LOG_ERROR_OUTPUT_PATHS
	errorOutputPaths := viper.GetStringSlice("LOG_ERROR_OUTPUT_PATHS")
	if len(errorOutputPaths) == 0 {
		if s := viper.GetString("LOG_ERROR_OUTPUT_PATHS"); s != "" {
			parts := strings.Split(s, ",")
			for i := range parts {
				parts[i] = strings.TrimSpace(parts[i])
			}
			errorOutputPaths = parts
		}
	}
	if len(errorOutputPaths) == 0 {
		errorOutputPaths = append([]string(nil), defaultErrorOutputPaths...)
	}
	Logger.ErrorOutputPaths = errorOutputPaths
}
