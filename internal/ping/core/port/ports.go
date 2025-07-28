package port

import (
	"context"
	"go-skeleton/internal/ping/core/domain"
)

// PingRepository defines the contract for ping repository operations
type PingRepository interface {
	Ping(ctx context.Context, resp *domain.Ping) error
}

// SvcContext holds all repository dependencies for the service layer
type SvcContext struct {
	Repo PingRepository
}

func NewServiceContext(
	pingRepository PingRepository,
) SvcContext {
	return SvcContext{
		Repo: pingRepository,
	}
}
