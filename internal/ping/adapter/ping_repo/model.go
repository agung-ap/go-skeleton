package pingrepo

import "go-skeleton/internal/ping/core/domain"

type PingResponse struct {
	Message string
}

func (r *PingResponse) ToDomain() domain.Ping {
	return domain.Ping{
		Message: r.Message,
	}
}
