package database

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
)

func NewRedisClient() *redis.Client {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	_, err := client.Ping(ctx).Result()

	if err != nil {
		panic(err)
	}

	return client
}
