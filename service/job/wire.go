//go:build wireinject
// +build wireinject

package service

import (
	"github.com/google/wire"
)

var JobServiceSet = wire.NewSet(
	NewJobService,
)
