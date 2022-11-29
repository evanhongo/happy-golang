//go:build wireinject
// +build wireinject

package job_queue

import (
	service "github.com/evanhongo/happy-golang/service/job"
	"github.com/google/wire"
)

var JobQueueSet = wire.NewSet(
	service.JobServiceSet,
	NewJobQueue,
)
