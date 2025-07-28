package pingrepo_test

import (
	"context"
	"go-skeleton/internal/ping/core/domain"
	"testing"

	"github.com/stretchr/testify/assert"

	pingrepo "go-skeleton/internal/ping/adapter/ping_repo"
)

func TestPingResponse_ToDomain(t *testing.T) {
	// Note: Testing the behavior through the repository since PingResponse is not exported
	// This tests the conversion indirectly
	ctx := context.Background()
	repo := pingrepo.NewPingRepository(nil, nil)
	var result domain.Ping

	// Act
	err := repo.Ping(ctx, &result)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "ping from repository", result.Message)
}
