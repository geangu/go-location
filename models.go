package main

import (
	"time"
)

// Location struct to manage location
type Location struct {
	Key       string    `json:"key"`
	Latitude  string    `json:"latitude"`
	Longitude string    `json:"longitude"`
	Date      time.Time `json:"date"`
}

// Create Location Object
func (l *Location) Create(key string, latitude string, longitude string) {
	l.Key = key
	l.Latitude = latitude
	l.Longitude = longitude
	l.Date = time.Now()

	l.saveRedisCache()
	l.saveLocationDB()
}

// GetRedisCache get last location registered in Redis
func (l *Location) GetRedisCache(key string) Location {
	redis := new(RedisHelper)
	return redis.GetLocation(key)
}

func (l *Location) saveLocationDB() {
	// TODO: implement me!!!
}

func (l *Location) saveRedisCache() {
	redis := new(RedisHelper)
	redis.SaveLocationRedis(l)
}
