package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

// RedisHelper struct
type RedisHelper struct{}

func (r *RedisHelper) getClient() (context.Context, *redis.Client) {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return ctx, client
}

// SaveLocationRedis method
func (r *RedisHelper) SaveLocationRedis(l *Location) {
	content, err := json.Marshal(l)
	if err != nil {
		panic(err)
	}

	ctx, client := r.getClient()
	err = client.Set(ctx, l.Key, string(content), 0).Err()
	if err != nil {
		panic(err)
	}
}

// GetLocation by key
func (r *RedisHelper) GetLocation(provider string) Location {
	var l Location
	ctx, client := r.getClient()
	val, err := client.Get(ctx, provider).Result()
	if err != nil {
		fmt.Println(err.Error())
		return l
	}

	err = json.Unmarshal([]byte(val), &l)
	if err != nil {
		fmt.Println(err.Error())
	}

	return l
}
