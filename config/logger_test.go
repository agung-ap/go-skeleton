package config

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func setupLoggerConfigForTest() {
	viper.Reset()
	viper.Set("LOG_LEVEL", "debug")
	viper.Set("LOG_DEVELOPMENT", "true")
	viper.Set("LOG_DISABLE_CALLER", "false")
	viper.Set("LOG_DISABLE_STACKTRACE", "false")
	viper.Set("LOG_ENCODING", "json")
	viper.Set("LOG_OUTPUT_PATHS", "stdout")
	viper.Set("LOG_ERROR_OUTPUT_PATHS", "stderr")
}

func TestInitLoggerConfig(t *testing.T) {
	setupLoggerConfigForTest()

	initLoggerConfig()

	assert.Equal(t, "debug", Logger.Level)
	assert.True(t, Logger.Development)
	assert.False(t, Logger.DisableCaller)
	assert.False(t, Logger.DisableStacktrace)
	assert.Equal(t, "json", Logger.Encoding)
	assert.Equal(t, []string{"stdout"}, Logger.OutputPaths)
	assert.Equal(t, []string{"stderr"}, Logger.ErrorOutputPaths)
}

func TestInitLoggerConfig_WithDefaults(t *testing.T) {
	viper.Reset() // Clear all values to test defaults

	initLoggerConfig()

	assert.Equal(t, "info", Logger.Level)
	assert.False(t, Logger.Development)
	assert.False(t, Logger.DisableCaller)
	assert.False(t, Logger.DisableStacktrace)
	assert.Equal(t, "json", Logger.Encoding)
	assert.Equal(t, []string{"stdout"}, Logger.OutputPaths)
	assert.Equal(t, []string{"stderr"}, Logger.ErrorOutputPaths)
}
