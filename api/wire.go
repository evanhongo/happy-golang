//go:build wireinject
// +build wireinject

package api

import (
	"github.com/evanhongo/happy-golang/api/route/auth"
	"github.com/evanhongo/happy-golang/api/route/job"
	"github.com/google/wire"
)

var HttpServerSet = wire.NewSet(
	auth.AuthHandlerSet,
	job.NewJobHandler,
	NewHttpServer,
)
