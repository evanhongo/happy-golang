package job_queue

import "github.com/evanhongo/happy-golang/entity"

type IJobQueue interface {
	Start() error
	SendJob(job *entity.JobRequest) (string, error)
	GetJobState(jobId string) (*entity.Job, error)
}
