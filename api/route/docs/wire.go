//go:build wireinject
// +build wireinject

package docs_route

import (
	"github.com/evanhongo/happy-golang/api"
	"github.com/google/wire"
)

func CreateRouter() (api.IRouter, error) {
	panic(wire.Build(
		NewDocsRouter,
	))
}
