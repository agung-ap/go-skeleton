package pingrepo_test

import (
	"context"
	"go-skeleton/internal/ping/core/domain"
	"testing"

	"github.com/stretchr/testify/assert"

	pingrepo "go-skeleton/internal/ping/adapter/ping_repo"
)

func TestPingRepository_Ping_Success_WithNilDependencies(t *testing.T) {
	// Test the repository with nil dependencies (no external calls)
	ctx := context.Background()
	repo := pingrepo.NewPingRepository(nil, nil)
	var result domain.Ping

	// Act
	err := repo.Ping(ctx, &result)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "ping from repository", result.Message)
}

func TestNewPingRepository(t *testing.T) {
	// Act
	repo := pingrepo.NewPingRepository(nil, nil)

	// Assert
	assert.NotNil(t, repo)
}
