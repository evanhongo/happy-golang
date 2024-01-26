package cache

import (
	"context"
	"time"
)

type ICache interface {
	Ping() error
	Get(ctx context.Context, key string)
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration)
	Del(ctx context.Context, keys ...string)
}
