package api

import (
	"fmt"
	"net/http"

	_ "github.com/evanhongo/happy-golang/api/docs"
	"github.com/evanhongo/happy-golang/api/route/auth"
	health "github.com/evanhongo/happy-golang/api/route/health"
	"github.com/evanhongo/happy-golang/api/route/job"
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
		// endpoint: /swagger/index.html
		url := ginSwagger.URL("/swagger/doc.json")
		engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}

	health.AddRoute(engine)

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
