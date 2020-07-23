package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	routes := app.Group("/location")
	{
		routes.GET("/:id", GetLocation)
		routes.POST("/:id", SaveLocation)
	}
	app.Run()
}
