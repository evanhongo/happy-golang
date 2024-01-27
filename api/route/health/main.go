package health_route

import (
	"github.com/evanhongo/happy-golang/api"
	"github.com/gin-gonic/gin"
)

type Router struct {
	rg      *gin.RouterGroup
	handler *PingHandler
}

func (r *Router) Register(g *gin.Engine) {
	r.rg = g.Group("/health")
	{
		// @Summary ping
		// @Description ping for test service alive or not
		// @Produce json
		// @Success 200 {string} string
		// @Failure 400 {string} string
		// @Router /ping [get]
		r.rg.GET("/ping", r.handler.Ping)
		r.rg.POST("/ping", r.handler.ValidateRequest, r.handler.Ping)
	}
}

func NewHealthRouter(handler *PingHandler) api.IRouter {
	return &Router{
		handler: handler,
	}
}
