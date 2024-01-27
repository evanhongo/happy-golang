package job_queue

import (
	service "github.com/evanhongo/happy-golang/service/job"
	"github.com/google/wire"
)

var CreateJobQueue = wire.NewSet(
	service.NewJobService,
	NewJobQueue,
)
