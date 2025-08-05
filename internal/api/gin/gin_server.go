package gin

import (
	"context"
	"crypto/tls"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpServer *http.Server
	engine     *gin.Engine
	enableTLS  bool
	certFile   string
	keyFile    string
}

func NewServer(cfg *GinHttpConfig) (*Server, error) {
	switch cfg.Mode {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	router := gin.Default()
	router.Use(MiddleWare())
	// 定义接口
	uri(router)

	httpServer := &http.Server{
		Addr:           fmt.Sprintf(":%d", cfg.Port),
		Handler:        router,
		ReadTimeout:    time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(cfg.WriteTimeout) * time.Second,
		MaxHeaderBytes: cfg.MaxHeaderBytes,
	}

	if cfg.EnableTLS {
		cert, err := tls.LoadX509KeyPair(cfg.TLSCertFile, cfg.TLSKeyFile)
		if err != nil {
			return nil, fmt.Errorf("failed to load TLS key pair: %w", err)
		}
		httpServer.TLSConfig = &tls.Config{
			Certificates: []tls.Certificate{cert},
		}
	}

	return &Server{
		httpServer: httpServer,
		engine:     router,
		enableTLS:  cfg.EnableTLS,
		certFile:   cfg.TLSCertFile,
		keyFile:    cfg.TLSKeyFile,
	}, nil
}

func (s *Server) Start() {
	go func() {
		addr := s.httpServer.Addr
		if s.enableTLS {
			slog.Info("gin server starting with TLS", "addr", addr)
			if err := s.httpServer.ListenAndServeTLS(s.certFile, s.keyFile); err != nil && err != http.ErrServerClosed {
				slog.Error("Could not start HTTPS server", "err", err)
			}
		} else {
			slog.Info("gin server starting", "addr", addr)
			if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				slog.Error("Could not start HTTP server", "err", err)
			}
		}
	}()
}

func (s *Server) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	slog.Info("Shutting down gin server...")
	if err := s.httpServer.Shutdown(ctx); err != nil {
		slog.Error("Could not shutdown http server", "err", err)
	}
	slog.Info("gin server shutdown successfully")
}
