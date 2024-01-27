// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package auth_route

import (
	"github.com/evanhongo/happy-golang/api"
	"github.com/evanhongo/happy-golang/service/auth"
)

// Injectors from wire.go:

func CreateRouter() (api.IRouter, error) {
	iAuthService := service.NewAuthService()
	authHandler := NewAuthHandler(iAuthService)
	iRouter := NewAuthRouter(authHandler)
	return iRouter, nil
}