package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitFromFile(t *testing.T) {
	Init()
	assert.NotNil(t, Server)
}

func TestInitForTest(t *testing.T) {
	InitForTest()
	assert.True(t, ConfigLoadedForTest)
}

func TestForMissingConfig(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotNil(t, r)
	}()
	InitForTest()
	// This should panic because DUMMY key doesn't exist
	mustGetString("DUMMY_NON_EXISTENT_KEY")
}

func TestForInvalidIntConfig(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotNil(t, r)
	}()
	InitForTest()
	// This should panic because LOG_LEVEL is not a valid integer
	mustGetInt("LOG_LEVEL")
}
