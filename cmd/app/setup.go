package app

import (
	"go-skeleton/config"
	dicontainer "go-skeleton/internal/common/di"
	"go-skeleton/internal/ping"
	"go-skeleton/pkg/cache"
	"go-skeleton/pkg/database"
	"go-skeleton/pkg/logger"

	"github.com/gin-gonic/gin"
)

var container dicontainer.Container

func Init() {
	config.Init()
	logger.Init(config.Logger)
	database.Init(config.Database)
	cache.Init(config.RedisCache)

	// Initialize dependency injection container
	container = dicontainer.NewContainer()
}

func SetupRouter() *gin.Engine {
	// Create global router with middleware and common settings
	router := NewGlobalRouter()

	// Initialize modules with dependency injection
	pingModule := ping.NewModule(container)

	// Register module routes
	pingModule.RegisterRoutes(router)

	return router
}

func ShutDown() {
	logger.Sync() // Flush any buffered logs
	database.CloseDB()
	cache.CloseCache()
}
