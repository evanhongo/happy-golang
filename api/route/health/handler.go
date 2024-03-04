package health_route

import (
	"net/http"

	"github.com/evanhongo/happy-golang/api/httputil"
	"github.com/evanhongo/happy-golang/pkg/schema"
	"github.com/gin-gonic/gin"
)

type PingRequest struct {
	Hello string `json:"hello"`
}

type PingHandler struct {
	schema schema.ISchema
}

func (handler *PingHandler) ValidateRequest(c *gin.Context) {
	var req PingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &httputil.HttpErrorBody{Error: err.Error()})
		c.Abort()
		return
	}

	if err := handler.schema.Parse(req); err != nil {
		c.JSON(http.StatusBadRequest, &httputil.HttpErrorBody{Error: err.Error()})
		c.Abort()
		return
	}

	c.Next()
}

// @Tags health
// @Summary ping
// @Description ping for test service alive or not
// @Produce json
// @Success 200 {string} string
// @Failure 400 {string} string
// @Router /ping [get]
func (handler *PingHandler) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func NewPingHandler(schema schema.ISchema) *PingHandler {
	return &PingHandler{
		schema,
	}
}
