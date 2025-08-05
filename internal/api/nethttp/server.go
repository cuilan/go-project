package nethttp

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

// Server 定义了一个 net/http 服务器。
type Server struct {
	httpServer *http.Server
	router     *http.ServeMux
}

// NewServer 创建一个新的 Server 实例。
func NewServer(cfg *NetHttpConfig) *Server {
	router := http.NewServeMux()

	httpServer := &http.Server{
		Addr:           fmt.Sprintf(":%d", cfg.Port),
		Handler:        router,
		ReadTimeout:    time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(cfg.WriteTimeout) * time.Second,
		MaxHeaderBytes: cfg.MaxHeaderBytes,
	}

	s := &Server{
		httpServer: httpServer,
		router:     router,
	}

	s.registerRoutes()
	return s
}

// Start 在一个新的 goroutine 中启动 HTTP 服务器。
func (s *Server) Start() {
	go func() {
		slog.Info("nethttp server starting", "addr", s.httpServer.Addr)
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Could not start http server", "err", err)
		}
	}()
}

// Stop 优雅地关闭 HTTP 服务器。
func (s *Server) Stop() error {
	slog.Info("nethttp server stopping")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.httpServer.Shutdown(ctx)
}
