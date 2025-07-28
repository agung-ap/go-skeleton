package rest

import (
	httpcommon "go-skeleton/internal/common/http"
	"go-skeleton/internal/ping/core/domain"
	"go-skeleton/internal/ping/core/service"
	"go-skeleton/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PingHandler struct {
	PingService *service.PingService
}

func NewPingHandler(
	pingService *service.PingService,
) *PingHandler {
	return &PingHandler{
		PingService: pingService,
	}
}

func (h *PingHandler) Ping(c *gin.Context) {
	logger.Info("Ping endpoint called",
		zap.String("remote_addr", c.ClientIP()),
		zap.String("user_agent", c.Request.UserAgent()),
	)

	var resp domain.Ping
	err := h.PingService.Ping(c, &resp)
	if err != nil {
		httpcommon.ResponseError(c, err)
		return
	}

	response := PingResponse{
		PingMessage: resp.Message,
	}

	httpcommon.ResponseSuccess(c, http.StatusOK, "success", response, nil)
}
