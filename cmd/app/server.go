package app

import (
	"context"
	"go-skeleton/config"
	"go-skeleton/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"go.uber.org/zap"
)

type Server struct {
	server *http.Server
}

func New() *Server {
	handler := SetupRouter()

	server := &Server{
		server: &http.Server{
			Addr:         ":" + strconv.Itoa(config.Server.Port),
			Handler:      handler,
			ReadTimeout:  config.Server.ReadTimeout,
			WriteTimeout: config.Server.WriteTimeout,
		},
	}

	logger.Info("Server created",
		zap.String("addr", server.server.Addr),
		zap.Duration("read_timeout", config.Server.ReadTimeout),
		zap.Duration("write_timeout", config.Server.WriteTimeout),
	)

	return server
}

func (s *Server) Start(ctx context.Context, cancel context.CancelFunc) {
	go s.waitForShutDown(ctx, cancel)

	logger.Info("Starting HTTP server", zap.String("addr", s.server.Addr))

	go func() {
		err := s.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Error("Server error", zap.Error(err))
			cancel()
		}
	}()
}

func (s *Server) waitForShutDown(ctx context.Context, cancel context.CancelFunc) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-stop
	logger.Info("Shutdown signal received, stopping server gracefully...")

	if err := s.server.Shutdown(ctx); err != nil {
		logger.Error("Server shutdown error", zap.Error(err))
	}

	logger.Info("Server stopped")
	cancel() // call the cancelFunc to close the shared interrupt channel between REST and gRPC and shutdown both servers
}
