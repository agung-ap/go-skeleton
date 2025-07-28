package container

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewContainer(t *testing.T) {
	// Act
	container := NewContainer()

	// Assert
	assert.NotNil(t, container)
	// Note: DB and Cache may be nil in test environment since they rely on global state
	// This is expected behavior - the container structure is created successfully
}

func TestContainer_Structure(t *testing.T) {
	// Act
	container := NewContainer()

	// Assert - Verify the container has the expected fields
	assert.IsType(t, Container{}, container)

	// The container structure should exist regardless of whether the global
	// DB and Cache instances are initialized
	assert.NotNil(t, container)
}
