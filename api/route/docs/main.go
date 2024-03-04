package docs_route

import (
	"github.com/evanhongo/happy-golang/api"
	_ "github.com/evanhongo/happy-golang/api/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
}

func (r *Router) Register(g *gin.Engine) {
	// swagger endpoint: /api-docs/index.html
	g.GET("/api-docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func NewDocsRouter() api.IRouter {
	return &Router{}
}
