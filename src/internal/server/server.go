package server

import (
	"context"
	"net/http"

	"github.com/avalonprod/gasstrem/src/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + cfg.HTTPConfig.Port,
			Handler:        handler,
			ReadTimeout:    cfg.HTTPConfig.ReadTimeout,
			WriteTimeout:   cfg.HTTPConfig.WriteTimeout,
			MaxHeaderBytes: cfg.HTTPConfig.MaxHeaderBytes << 25,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
