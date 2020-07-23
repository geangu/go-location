package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// RedisHelper struct
type RedisHelper struct{}

func (r *RedisHelper) getClient() (context.Context, *redis.Client) {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return ctx, client
}

// SaveLocationRedis method
func (r *RedisHelper) SaveLocationRedis(l *Location) {
	content, err := json.Marshal(l)
	if err != nil {
		log.Fatal(err)
	}

	ctx, client := r.getClient()
	err = client.Set(ctx, l.Key, string(content), 0).Err()
	if err != nil {
		log.Fatal(err)
	}
}

// GetLocation by key
func (r *RedisHelper) GetLocation(provider string) Location {
	var l Location
	ctx, client := r.getClient()
	val, err := client.Get(ctx, provider).Result()
	if err != nil {
		log.Fatal(err.Error())
		return l
	}

	err = json.Unmarshal([]byte(val), &l)
	if err != nil {
		log.Fatal(err.Error())
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
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("location_db").Collection("locations")
	collection.InsertOne(context.TODO(), location)

	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
}
