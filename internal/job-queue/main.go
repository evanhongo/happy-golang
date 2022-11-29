package job_queue

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	machineryLog "github.com/RichardKnop/machinery/v1/log"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/evanhongo/happy-golang/entity"
	"github.com/evanhongo/happy-golang/internal/env"
	"github.com/evanhongo/happy-golang/pkg/logger"
	service "github.com/evanhongo/happy-golang/service/job"
	"github.com/google/uuid"
)

type JobQueue struct {
	server *machinery.Server
}

func (queue *JobQueue) Start() error {
	env := env.GetEnv()
	worker := queue.server.NewWorker("my-worker", env.JOB_QUEUE_WORKER_NUM)
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
	env := env.GetEnv()
	redisConfig := &config.RedisConfig{}
	machineryConfig := &config.Config{
		Broker:          env.JOB_QUEUE_BROKER,
		ResultBackend:   env.JOB_QUEUE_RESULT_BACKEND,
		ResultsExpireIn: env.JOB_QUEUE_RESULT_EXPIRE_IN,
		Redis:           redisConfig,
		DefaultQueue:    env.JOB_QUEUE_DEFAULT_QUEUE,
	}

	server, err := machinery.NewServer(machineryConfig)
	if err != nil {
		return nil, err
	}

	server.RegisterTasks(map[string]interface{}{
		COMPRESS_IMAGE: service.CompressImage,
	})

	logLevel := env.JOB_QUEUE_LOG_LEVEL
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
			machineryLog.SetDebug(infoLogger)
		case "WARNING":
			warningLogger := NewLogger("warning")
			machineryLog.SetDebug(warningLogger)
		case "ERROR":
			errorLogger := NewLogger("error")
			machineryLog.SetDebug(errorLogger)
		case "FATAL":
			fatalLogger := NewLogger("fatal")
			machineryLog.SetDebug(fatalLogger)
		}
		if l == logLevel {
			break
		}
	}

	return &JobQueue{
		server,
	}, nil
}
