package app

import (
	"go-skeleton/config"
	"go-skeleton/pkg/logger"

	"github.com/gin-gonic/gin"
)

// NewGlobalRouter creates and configures the global router with middleware and common settings
func NewGlobalRouter() *gin.Engine {
	// Set Gin mode based on log level for better performance in production
	if config.Logger.Level == "debug" || config.Logger.Development {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	// Add global middleware
	router.Use(gin.Recovery())             // Equivalent to Chi's Recoverer
	router.Use(logger.LoggingMiddleware()) // Our custom logging middleware

	// Static file serving for docs
	router.Static("/docs", config.App.DocsPath)

	return router
}
