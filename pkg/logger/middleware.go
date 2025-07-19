package logger

import (
	httpcommon "go-skeleton/internal/common/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GinLoggingMiddleware logs HTTP requests and responses for Gin
func GinLoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Log request
		GetLogger().Info("HTTP Request",
			zap.String("method", c.Request.Method),
			zap.String("url", c.Request.URL.String()),
			zap.String("remote_addr", c.ClientIP()),
			zap.String("user_agent", c.Request.UserAgent()),
			zap.String("request_id", c.GetHeader(httpcommon.HeaderRequestID.String())),
		)

		// Process request
		c.Next()

		// Log response
		duration := time.Since(start)
		GetLogger().Info("HTTP Response",
			zap.String("method", c.Request.Method),
			zap.String("url", c.Request.URL.String()),
			zap.Int("status_code", c.Writer.Status()),
			zap.Duration("duration", duration),
			zap.String("request_id", c.GetHeader("X-Request-ID")),
		)
	}
}

// LoggingMiddleware logs HTTP requests and responses (kept for backward compatibility)
func LoggingMiddleware() gin.HandlerFunc {
	return GinLoggingMiddleware()
}
