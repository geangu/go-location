package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type saveLocationRequest struct {
	Longitude string `json:"longitude" binding:"required"`
	Latitude  string `json:"latitude" binding:"required"`
}

// GetLocation get last location from Redis
func GetLocation(c *gin.Context) {
	id := c.Param("id")
	location := new(Location)
	result := location.GetRedisCache(id)
	if result.Key == "" {
		status := http.StatusNotFound
		c.JSON(status, http.StatusText(status))
		return
	}
	c.JSON(http.StatusOK, result)
}

// SaveLocation save the last location and Update Redis item
func SaveLocation(c *gin.Context) {
	id := c.Param("id")
	var data saveLocationRequest
	c.BindJSON(&data)

	location := new(Location)
	location.Create(
		id,
		data.Latitude,
		data.Longitude,
	)
	c.JSON(http.StatusOK, location)
}
