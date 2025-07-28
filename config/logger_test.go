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

func TestGetStringOrDefault(t *testing.T) {
	viper.Reset()
	viper.Set("EXISTING_KEY", "existing_value")

	// Test with existing key
	result := getStringOrDefault("EXISTING_KEY", "default_value")
	assert.Equal(t, "existing_value", result)

	// Test with non-existing key
	result = getStringOrDefault("NON_EXISTING_KEY", "default_value")
	assert.Equal(t, "default_value", result)
}

func TestGetBoolOrDefault(t *testing.T) {
	viper.Reset()
	viper.Set("BOOL_TRUE", "true")
	viper.Set("BOOL_FALSE", "false")

	// Test with existing true value
	result := getBoolOrDefault("BOOL_TRUE", false)
	assert.True(t, result)

	// Test with existing false value
	result = getBoolOrDefault("BOOL_FALSE", true)
	assert.False(t, result)

	// Test with non-existing key
	result = getBoolOrDefault("NON_EXISTING_KEY", true)
	assert.True(t, result)
}

func TestGetStringSliceOrDefault(t *testing.T) {
	viper.Reset()
	viper.Set("EXISTING_KEY", "single_value")

	// Test with existing key
	result := getStringSliceOrDefault("EXISTING_KEY", []string{"default1", "default2"})
	assert.Equal(t, []string{"single_value"}, result)

	// Test with non-existing key
	result = getStringSliceOrDefault("NON_EXISTING_KEY", []string{"default1", "default2"})
	assert.Equal(t, []string{"default1", "default2"}, result)
}
