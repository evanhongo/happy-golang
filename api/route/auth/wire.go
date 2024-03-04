//go:build wireinject
// +build wireinject

package auth_route

import (
	"github.com/evanhongo/happy-golang/api"
	"github.com/evanhongo/happy-golang/pkg/db"
	service "github.com/evanhongo/happy-golang/service/auth"
	"github.com/google/wire"
)

func CreateRouter() (api.IRouter, error) {
	panic(wire.Build(
		db.NewDb,
		service.NewAuthService,
		NewAuthHandler,
		NewAuthRouter,
	))
}
