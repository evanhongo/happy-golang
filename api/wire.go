package api

import (
	"github.com/evanhongo/happy-golang/internal/env"
	"github.com/google/wire"
)

func DefaultPort() string {
	return env.GetEnv().PORT
}

var CreateServer = wire.NewSet(
	DefaultPort,
	NewServer,
)
