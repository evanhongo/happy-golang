//go:build wireinject
// +build wireinject

package main

import (
	"github.com/evanhongo/happy-golang/api"
	job_queue "github.com/evanhongo/happy-golang/internal/job-queue"
	"github.com/google/wire"
)

func CreateServer() (*Server, error) {
	panic(wire.Build(
		api.HttpServerSet,
		job_queue.JobQueueSet,
		NewServer,
	))
}
