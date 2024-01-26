package health_route

import (
	"github.com/gin-gonic/gin"
)

func AddRoute(r *gin.Engine) {
	handler := &PingHandler{
		schema: NewPingRequestSchema(),
	}
	router := r.Group("/health")
	{
		// @Summary ping
		// @Description ping for test service alive or not
		// @Produce json
		// @Success 200 {string} string
		// @Failure 400 {string} string
		// @Router /ping [get]
		router.GET("/ping", handler.Ping)
		router.POST("/ping", handler.ValidateRequest, handler.Ping)
	}
}
