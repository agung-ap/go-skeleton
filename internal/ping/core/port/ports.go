package port

import (
	pingrepo "go-skeleton/internal/ping/adapter/ping_repo"
)

// Context holds all repository dependencies for the service layer
type SvcContext struct {
	Repo pingrepo.PingRepository
}

func NewServiceContext(
	pingRepository pingrepo.PingRepository,
) SvcContext {
	return SvcContext{
		Repo: pingRepository,
	}
}
