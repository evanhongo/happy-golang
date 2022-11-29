package cache

import (
	"context"
	"time"
)

//go:generate mockgen -source interface.go -destination ../../mock/mock_cache.go -package mock

type ICache interface {
	Ping() error
	Get(ctx context.Context, key string)
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration)
	Del(ctx context.Context, keys ...string)
}
