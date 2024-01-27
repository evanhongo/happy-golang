package health_route

import (
	"encoding/json"
	"net/http"

	api "github.com/evanhongo/happy-golang/api/httputil"
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
	var data any
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, &api.HttpErrorBody{Code: string(api.INVALID_REQUEST), Error: err.Error()})
		c.Abort()
		return
	}

	jsonStr, err := json.Marshal(data)
	if err != nil {
		c.JSON(400, &api.HttpErrorBody{Code: string(api.INVALID_REQUEST), Error: err.Error()})
		c.Abort()
		return
	}

	var structData PingRequest
	if err := json.Unmarshal(jsonStr, &structData); err != nil {
		c.JSON(400, &api.HttpErrorBody{Code: string(api.INVALID_REQUEST), Error: err.Error()})
		c.Abort()
		return
	}

	if err := handler.schema.Parse(structData); err != nil {
		c.JSON(400, &api.HttpErrorBody{Code: string(api.INVALID_REQUEST), Error: err.Error()})
		c.Abort()
		return
	}

	c.Next()
}

func (handler *PingHandler) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func NewPingHandler(schema schema.ISchema) *PingHandler {
	return &PingHandler{
		schema,
	}
}
