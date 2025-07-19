package ping

import (
	dicontainer "go-skeleton/internal/common/container"
	pingrepo "go-skeleton/internal/ping/adapter/ping_repo"
	restHandl "go-skeleton/internal/ping/adapter/rest"
	"go-skeleton/internal/ping/core/port"
	"go-skeleton/internal/ping/core/service"

	"github.com/gin-gonic/gin"
)

type Module struct {
	Service     service.PingService
	RestHandler restHandl.PingHandler
	Router      restHandl.Router
}

func NewModule(container dicontainer.Container) Module {
	// Inject database container to repository
	pingRepo := pingrepo.CreateEnhancedRepository(container.DB, container.Cache)

	// Create service context to group repositories
	serviceCtx := port.NewServiceContext(pingRepo)

	// Create service with injected context
	svc := service.NewPingService(serviceCtx)

	// Create handlers and router
	restHandler := restHandl.NewPingHandler(svc)
	router := restHandl.NewRouter(restHandler)

	return Module{
		Service:     svc,
		RestHandler: restHandler,
		Router:      router,
	}
}

// RegisterRoutes registers all ping routes to the provided router
func (m Module) RegisterRoutes(router *gin.Engine) {
	m.Router.RegisterPingRoutes(router)
}
