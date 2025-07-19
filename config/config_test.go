package config_test

import (
	"go-skeleton/config"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestInitFromFile(t *testing.T) {
	config.Init()
	assert.NotNil(t, config.Server)
}

func TestInitForTest(t *testing.T) {
	config.InitForTest()
	assert.True(t, config.ConfigLoadedForTest)
}

func TestForMissingConfig(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotNil(t, r)
	}()
	config.InitForTest()
	viper.GetString("DUMMY")
}

func TestForInvalidIntConfig(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotNil(t, r)
	}()
	config.InitForTest()
	viper.GetString("LOG_LEVEL")
}
