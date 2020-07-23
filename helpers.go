package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// RedisHelper struct
type RedisHelper struct{}

func (r *RedisHelper) getClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}

// SaveLocationRedis method
func (r *RedisHelper) SaveLocationRedis(l *Location) {
	content, err := json.Marshal(l)
	if err != nil {
		fmt.Println(err)
	}

	client := r.getClient()
	err = client.Set(l.Key, string(content), 0).Err()
	if err != nil {
		fmt.Println(err)
	}
}

// GetLocation by key
func (r *RedisHelper) GetLocation(provider string) Location {
	var l Location
	client := r.getClient()
	val, err := client.Get(provider).Result()
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

// MongoHelper struct
type MongoHelper struct{}

// InsertLocation item to collection Item
func (m *MongoHelper) InsertLocation(location Location) {

	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println(err)
	}

	collection := client.Database("location_db").Collection("locations")
	collection.InsertOne(context.TODO(), location)

	err = client.Disconnect(context.TODO())

	if err != nil {
		fmt.Println(err)
	}
}
