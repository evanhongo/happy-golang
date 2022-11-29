package api

import (
	"fmt"
	"net/http"

	_ "github.com/evanhongo/happy-golang/api/docs"
	"github.com/evanhongo/happy-golang/api/handlers/auth"
	"github.com/evanhongo/happy-golang/api/handlers/job"
	"github.com/evanhongo/happy-golang/internal/env"
	pb "github.com/evanhongo/happy-golang/rpc/job"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go Service Demo
// @version 1.0
// @description Swagger API.
// @contact.email evan@example.com
func NewHttpServer(jobHandler *job.JobHandler, authHandler *auth.AuthHandler) *http.Server {
	env := env.GetEnv()
	if env.ENVIRONMENT == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()

	if env.ENVIRONMENT != "production" {
		// swagger
		url := ginSwagger.URL("/swagger/doc.json")
		engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}

	engine.GET("/ping", Ping)
	engine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	engine.GET("/auth", authHandler.RetrieveAuthorizationCode)
	engine.GET("/auth/callback", authHandler.RetrieveAccessToken)

	twirpHandler := pb.NewJobServiceServer(jobHandler)
	rpcPath := fmt.Sprintf("%s:method", twirpHandler.PathPrefix())
	engine.POST(rpcPath, jobHandler.CheckRPCMethod, gin.WrapH(twirpHandler))

	engine.GET("/job/:jobId", jobHandler.GetJobState)

	return &http.Server{
		Addr:    ":" + env.PORT,
		Handler: engine,
	}
}

// @Summary ping
// @Description ping for test service alive or not
// @Produce json
// @Success 200 {string} string
// @Failure 400 {string} string
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
