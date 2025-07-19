package service

import "go-skeleton/internal/ping/core/port"

// PingService implements business logic with repository access through context
type PingService struct {
	svcCtx port.SvcContext
}

// NewPingService creates a service with injected service context
func NewPingService(svcCtx port.SvcContext) PingService {
	return PingService{
		svcCtx: svcCtx,
	}
}

// Ping returns a ping response through the service context
func (s PingService) Ping() string {
	return s.svcCtx.Repo.Ping()
}
