package health_route

import (
	"github.com/evanhongo/happy-golang/api"
	"github.com/gin-gonic/gin"
)

type Router struct {
	handler *PingHandler
}

func (r *Router) Register(g *gin.Engine) {
	rg := g.Group("/health")
	{
		rg.GET("/ping", r.handler.Ping)
		rg.POST("/ping", r.handler.ValidateRequest, r.handler.Ping)
	}
}

func NewHealthRouter(handler *PingHandler) api.IRouter {
	return &Router{
		handler: handler,
	}
}
