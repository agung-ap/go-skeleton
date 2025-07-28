package port

import (
	pingrepo "go-skeleton/internal/ping/adapter/ping_repo"
)

type PingRepository = pingrepo.PingRepository

// SvcContext holds all repository dependencies for the service layer
type SvcContext struct {
	Repo *PingRepository
}

func NewServiceContext(
	pingRepository *pingrepo.PingRepository,
) *SvcContext {
	return &SvcContext{
		Repo: pingRepository,
	}
}
