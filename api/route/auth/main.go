package auth_route

import (
	"github.com/evanhongo/happy-golang/api"
	"github.com/gin-gonic/gin"
)

type Router struct {
	handler *AuthHandler
}

func (r *Router) Register(g *gin.Engine) {
	rg := g.Group("/auth")
	{
		rg2 := rg.Group("/email")
		{
			rg2.POST("/register", r.handler.RegisterByEmail)
			rg2.POST("/login", r.handler.LoginByEmail)
		}

		rg3 := rg.Group("/google")
		{
			rg3.POST("/login", r.handler.RetrieveAuthorizationCode)
			rg3.POST("/callback", r.handler.RetrieveAccessToken)
		}

	}
}

func NewAuthRouter(handler *AuthHandler) api.IRouter {
	return &Router{
		handler: handler,
	}
}
