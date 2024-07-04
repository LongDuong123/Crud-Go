package redis

import (
	"github.com/go-redis/redis/v8"
)

type databaseRedis struct {
	redis *redis.Client
}

func ConnnectRedis() *databaseRedis {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-container:6379",
		Password: "",
		DB:       0,
	})
	return &databaseRedis{redis: rdb}
}
