package auth_route

import (
	"github.com/evanhongo/happy-golang/api"
	"github.com/gin-gonic/gin"
)

type Router struct {
	rg      *gin.RouterGroup
	handler *AuthHandler
}

func (r *Router) Register(g *gin.Engine) {
	r.rg = g.Group("/auth")
	{
		r.rg.GET("/", r.handler.RetrieveAuthorizationCode)
		r.rg.GET("/callback", r.handler.RetrieveAccessToken)
	}
}

func NewAuthRouter(handler *AuthHandler) api.IRouter {
	return &Router{
		handler: handler,
	}
}
