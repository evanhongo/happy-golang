package cache

import (
	"context"
	"time"
)

type Cache struct {
}

func (cache *Cache) Ping() error {
	return nil
}

func (cache *Cache) Get(ctx context.Context, key string) {
	return
}

func (cache *Cache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) {
	return
}

func (cache *Cache) Del(ctx context.Context, keys ...string) {
	return
}

func NewCache() (ICache, error) {
	return &Cache{}, nil
}
