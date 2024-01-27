//go:build wireinject
// +build wireinject

package job_route

import (
	"github.com/evanhongo/happy-golang/api"
	job_queue "github.com/evanhongo/happy-golang/pkg/job_queue"
	"github.com/google/wire"
)

func CreateRouter() (api.IRouter, error) {
	panic(wire.Build(
		job_queue.CreateJobQueue,
		NewJobHandler,
		NewJobRouter,
	))
}
