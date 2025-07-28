package service

import (
	"context"
	"go-skeleton/internal/ping/core/domain"
	"go-skeleton/internal/ping/core/port"
)

// PingService implements business logic with repository access through context
type PingService struct {
	svcCtx *port.SvcContext
}

// NewPingService creates a service with injected service context
func NewPingService(svcCtx *port.SvcContext) PingService {
	return PingService{
		svcCtx: svcCtx,
	}
}

// Ping returns a ping response through the service context
func (s *PingService) Ping(ctx context.Context, resp *domain.Ping) error {
	return s.svcCtx.Repo.Ping(ctx, resp)
}
