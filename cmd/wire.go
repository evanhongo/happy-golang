//go:build wireinject
// +build wireinject

package main

import (
	"github.com/evanhongo/happy-golang/api"

	job_queue "github.com/evanhongo/happy-golang/pkg/job_queue"
	"github.com/google/wire"
)

func CreateCmd() (*Cmd, error) {
	panic(wire.Build(
		api.CreateServer,
		job_queue.CreateJobQueue,
		NewCmd,
	))
}
