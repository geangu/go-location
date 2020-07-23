package main

import (
	"time"
)

// Location struct to manage location
type Location struct {
	Provider  string    `json:"provider"`
	Latitude  string    `json:"latitude"`
	Longitude string    `json:"longitude"`
	Date      time.Time `json:"date"`
}

// Create Location Object
func (l *Location) Create(provider string, latitude string, longitude string) {
	l.Provider = provider
	l.Latitude = latitude
	l.Longitude = longitude
	l.Date = time.Now()

	l.saveRedisCache()
	l.saveLocationDB()
}

// GetRedisCache get last location registered in Redis
func (l *Location) GetRedisCache(provider string) string {
	redis := new(RedisHelper)
	return redis.GetProviderLocation(provider)
}

func (l *Location) saveLocationDB() {
	// TODO: implement me!!!
}

func (l *Location) saveRedisCache() {
	redis := new(RedisHelper)
	redis.SaveLocationRedis(l)
}
