package api

import (
	"context"
	"flow-run/internal/flowrun/config"
	"flow-run/internal/lib/logger"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	Middleware interface {
		Handler() gin.HandlerFunc
	}

	Handler interface {
		Group() string
		Method() string
		Path() string
		Handle(*gin.Context)
	}
)

type Server struct {
	httpServer http.Server
}

func NewServer(middlewares []Middleware, handlers []Handler, cfg *config.Config) *Server {
	r := gin.New()

	for _, m := range middlewares {
		r.Use(m.Handler())
	}

	for _, h := range handlers {
		r.Group(h.Group()).Handle(h.Method(), h.Path(), h.Handle)
	}

	return &Server{
		httpServer: http.Server{
			Addr:    fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.ServerPort),
			Handler: r.Handler(),
		},
	}
}

func (s *Server) Start(ctx context.Context) error {
	ch := make(chan error)

	go func() {
		defer close(ch)

		if err := s.httpServer.ListenAndServe(); err != http.ErrServerClosed {
			ch <- err
			logger.Log.WithError(err).Error("Failed to start server")
		}
	}()

	select {
	case <-ctx.Done():
		logger.Log.Info("web server started")
	case err := <-ch:
		return err
	}
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
