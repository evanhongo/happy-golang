package docs_route

import (
	"github.com/evanhongo/happy-golang/api"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
}

func (r *Router) Register(g *gin.Engine) {
	// swagger
	// endpoint: /swagger/index.html
	url := ginSwagger.URL("/swagger/doc.json")
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func NewDocsRouter() api.IRouter {
	return &Router{}
}
