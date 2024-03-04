package job_route

import (
	"fmt"

	"github.com/evanhongo/happy-golang/api"
	pb "github.com/evanhongo/happy-golang/rpc/job"
	"github.com/gin-gonic/gin"
)

type Router struct {
	handler *JobHandler
}

func (r *Router) Register(g *gin.Engine) {
	twirpHandler := pb.NewJobServiceServer(r.handler)
	rpcPath := fmt.Sprintf("%s:method", twirpHandler.PathPrefix())
	g.POST(rpcPath, r.handler.CheckRPCMethod, gin.WrapH(twirpHandler))
	rg := g.Group("/job")
	{
		rg.GET("/:jobId", r.handler.GetJobState)
	}
}

func NewJobRouter(handler *JobHandler) api.IRouter {
	return &Router{
		handler: handler,
	}
}
