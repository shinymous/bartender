package database

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func NewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS"),
		Password: "",
		DB:       0,
	})
	return rdb.WithContext(ctx)
}
