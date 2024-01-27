package api

import (
	"context"
	"net/http"

	"github.com/evanhongo/happy-golang/internal/env"
	"github.com/gin-gonic/gin"
)

// @title Go Service Demo
// @version 1.0
// @description Swagger API.
// @contact.email evan@example.com

type IRouter interface {
	Register(g *gin.Engine)
}

type Server struct {
	port string
	g    *gin.Engine
	h    *http.Server
}

func (s *Server) Init() {
	env := env.GetEnv()
	if env.ENVIRONMENT == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	s.g = gin.New()

}

func (s *Server) RegisterRouter(r IRouter) {
	r.Register(s.g)
}

func (s *Server) Start() error {
	s.h = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.g,
	}

	if err := s.h.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.h.Shutdown(ctx)
}

func NewServer(port string) *Server {
	return &Server{
		port: port,
	}
}
