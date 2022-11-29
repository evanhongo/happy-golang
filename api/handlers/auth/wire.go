package auth

import (
	service "github.com/evanhongo/happy-golang/service/auth"
	"github.com/google/wire"
)

var AuthHandlerSet = wire.NewSet(
	service.NewAuthService,
	NewAuthHandler,
)
