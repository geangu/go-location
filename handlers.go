package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProviderLocation get last provider location from Redis
func GetProviderLocation(c *gin.Context) {
	location := new(Location)
	result := location.GetRedisCache("1")
	c.JSON(http.StatusOK, result)
}

// SaveProviderLocation save the last location for the provider and Update Redis item
func SaveProviderLocation(c *gin.Context) {
	location := new(Location)
	location.Create("1", "54.01", "32.04")
	c.JSON(http.StatusOK, location)
}
