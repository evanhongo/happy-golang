package api

import (
	"github.com/evanhongo/happy-golang/config"
	"github.com/google/wire"
)

func DefaultPort() string {
	return config.GetConfig().PORT
}

var CreateServer = wire.NewSet(
	DefaultPort,
	NewServer,
)
