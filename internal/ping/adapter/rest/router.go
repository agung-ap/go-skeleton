package rest

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	handler PingHandler
}

func NewRouter(handler PingHandler) Router {
	return Router{
		handler: handler,
	}
}

// RegisterPingRoutes registers ping-specific routes to the provided router
func (r Router) RegisterPingRoutes(router *gin.Engine) {
	router.GET("/ping", r.handler.Ping)
}
