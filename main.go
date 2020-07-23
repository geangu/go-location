package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	app := gin.Default()
	routes := app.Group("/location")
	{
		routes.GET("/provider/:id", GetProviderLocation)
		routes.POST("/provider/:id", SaveProviderLocation)
	}
	app.Run()
}
