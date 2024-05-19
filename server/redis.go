package server

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

func NewRedisConnection(port, password string) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:" + port,
		Password: password,
		DB: 0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}