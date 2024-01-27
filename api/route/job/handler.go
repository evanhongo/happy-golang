package job_route

import (
	"context"
	"net/http"

	api "github.com/evanhongo/happy-golang/api/httputil"
	job_queue "github.com/evanhongo/happy-golang/pkg/job_queue"
	"github.com/evanhongo/happy-golang/pkg/logger"
	pb "github.com/evanhongo/happy-golang/rpc/job"
	"github.com/gin-gonic/gin"
)

type JobHandler struct {
	jobQueue job_queue.IJobQueue
}

// @Tags job
// @Summary get state of job
// @Description get job id, name, state, result, error
// @Produce json
// @Param jobId path string true "job id"
// @Success 200 {object} entity.Job
// @Failure 400 {object} api.HttpErrorBody
// @Router /state/{jobId} [get]
func (handler *JobHandler) GetJobState(c *gin.Context) {
	jobId := c.Param("jobId")
	job, err := handler.jobQueue.GetJobState(jobId)
	if err != nil {
		c.JSON(http.StatusNotFound, &api.HttpErrorBody{Code: string(api.NOT_FOUND), Error: err.Error()})
	} else {
		c.JSON(http.StatusOK, job)
	}
}

func (handler *JobHandler) CompressImage(ctx context.Context, req *pb.CompressImageReq) (*pb.CompressImageResp, error) {
	jobId, err := handler.jobQueue.SendJob(job_queue.NewJobRequest(job_queue.COMPRESS_IMAGE, req.Data))
	if err != nil {
		return nil, err
	}
	return &pb.CompressImageResp{JobId: jobId}, nil
}

func (handler *JobHandler) CheckRPCMethod(c *gin.Context) {
	method := c.Param("method")
	if isSupportedRPCMethod(method) {
		c.Next()
	} else {
		logger.Error("Unknown rpc method")
		c.AbortWithStatusJSON(http.StatusBadRequest, &api.HttpErrorBody{Code: string(api.INVALID_REQUEST), Error: "Unknown rpc method"})
	}
}

func isSupportedRPCMethod(method string) bool {
	supportedMethods := []string{"CompressImage"}
	for _, v := range supportedMethods {
		if v == method {
			return true
		}
	}

	return false
}

func NewJobHandler(jobQueue job_queue.IJobQueue) *JobHandler {
	return &JobHandler{
		jobQueue,
	}
}
