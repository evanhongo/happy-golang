package job_queue

import (
	"github.com/RichardKnop/machinery/v1"
	redisCfg "github.com/RichardKnop/machinery/v1/config"
	machineryLog "github.com/RichardKnop/machinery/v1/log"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/evanhongo/happy-golang/config"
	"github.com/evanhongo/happy-golang/entity"
	"github.com/evanhongo/happy-golang/pkg/logger"
	service "github.com/evanhongo/happy-golang/service/job"
	"github.com/google/uuid"
)

type JobQueue struct {
	server *machinery.Server
}

func (queue *JobQueue) Start() error {
	cfg := config.GetConfig()
	worker := queue.server.NewWorker("my-worker", cfg.JOB_QUEUE_WORKER_NUM)
	worker.SetErrorHandler(func(err error) {
		logger.Error(err)
	})
	return worker.Launch()
}

func (queue *JobQueue) SendJob(job *entity.JobRequest) (string, error) {
	signature, err := newSignature(job)
	if err != nil {
		logger.Error(err)
		return "", err
	}

	result, err := queue.server.SendTask(signature)
	if err != nil {
		logger.Error(err)
		return "", err
	}
	jobId := result.GetState().TaskUUID

	return jobId, nil
}

func (queue *JobQueue) GetJobState(jobId string) (*entity.Job, error) {
	state, err := queue.server.GetBackend().GetState(jobId)
	if err != nil {
		logger.Error(jobId, err)
		return nil, err
	}

	return newJob(state), nil
}

func NewJobRequest(name string, data interface{}) *entity.JobRequest {
	return &entity.JobRequest{
		Name: name,
		Data: data,
	}
}

func newJob(state *tasks.TaskState) *entity.Job {
	job := &entity.Job{}
	job.Id = state.TaskUUID
	job.Name = state.TaskName
	job.State = state.State
	job.Error = state.Error
	if state.State == tasks.StateSuccess && len(state.Results) > 0 {
		job.Result = state.Results[0].Value
	} else {
		job.Result = nil
	}

	return job
}

func newSignature(job *entity.JobRequest) (*tasks.Signature, error) {
	jobId := uuid.New().String()
	return &tasks.Signature{
		UUID: jobId,
		Name: job.Name,
		Args: []tasks.Arg{
			{Type: "[]uint8", Value: job.Data},
		},
	}, nil
}

func NewJobQueue(service service.IJobService) (IJobQueue, error) {
	cfg := config.GetConfig()
	redisConfig := &redisCfg.RedisConfig{}
	machineryConfig := &redisCfg.Config{
		Broker:          cfg.JOB_QUEUE_BROKER,
		ResultBackend:   cfg.JOB_QUEUE_RESULT_BACKEND,
		ResultsExpireIn: cfg.JOB_QUEUE_RESULT_EXPIRE_IN,
		Redis:           redisConfig,
		DefaultQueue:    cfg.JOB_QUEUE_DEFAULT_QUEUE,
	}

	server, err := machinery.NewServer(machineryConfig)
	if err != nil {
		return nil, err
	}

	server.RegisterTasks(map[string]interface{}{
		COMPRESS_IMAGE: service.CompressImage,
	})

	logLevel := cfg.JOB_QUEUE_LOG_LEVEL
	logLevels := []string{
		"FATAL",
		"ERROR",
		"WARNING",
		"INFO",
		"DEBUG",
	}
	for _, l := range logLevels {
		switch l {
		case "DEBUG":
			debugLogger := NewLogger("debug")
			machineryLog.SetDebug(debugLogger)
		case "INFO":
			infoLogger := NewLogger("info")
			machineryLog.SetInfo(infoLogger)
		case "WARNING":
			warningLogger := NewLogger("warning")
			machineryLog.SetWarning(warningLogger)
		case "ERROR":
			errorLogger := NewLogger("error")
			machineryLog.SetError(errorLogger)
		case "FATAL":
			fatalLogger := NewLogger("fatal")
			machineryLog.SetFatal(fatalLogger)
		}
		if l == logLevel {
			break
		}
	}

	return &JobQueue{
		server,
	}, nil
}
