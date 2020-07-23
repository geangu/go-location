package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"
)

// RedisHelper struct
type RedisHelper struct{}

// SaveLocationRedis method
func (r *RedisHelper) SaveLocationRedis(l *Location) {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err := client.Set(ctx, l.Provider, l.Provider, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Redis updated")
}

// GetProviderLocation method by provider id
func (r *RedisHelper) GetProviderLocation(provider string) string {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	val, err := client.Get(ctx, provider).Result()
	if err != nil {
		fmt.Println(err)
	}
	return val
}
