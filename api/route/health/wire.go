//go:build wireinject
// +build wireinject

package health_route

import (
	"github.com/evanhongo/happy-golang/api"
	"github.com/google/wire"
)

func CreateRouter() (api.IRouter, error) {
	panic(wire.Build(
		NewPingRequestSchema,
		NewPingHandler,
		NewHealthRouter,
	))
}
