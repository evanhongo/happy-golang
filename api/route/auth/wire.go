//go:build wireinject
// +build wireinject

package auth_route

import (
	"github.com/evanhongo/happy-golang/api"
	service "github.com/evanhongo/happy-golang/service/auth"
	"github.com/google/wire"
)

func CreateRouter() (api.IRouter, error) {
	panic(wire.Build(
		service.NewAuthService,
		NewAuthHandler,
		NewAuthRouter,
	))
}
