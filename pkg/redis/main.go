package redis

import (
	"github.com/evanhongo/happy-golang/config"
	"github.com/go-redis/redis"
)

func NewRedis() (*redis.Client, error) {
	cfg := config.GetConfig()
	options := redis.Options{
		Network: "tcp",
		Addr:    cfg.REDIS_ENDPOINT,
	}
	redisClient := redis.NewClient(&options)
	_, err := redisClient.Ping().Result()
	if err != nil {
		return nil, err
	}
	return redisClient, nil
}
