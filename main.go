package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ricardolonga/golang/controller"
)

func main() {
	router := gin.Default()

	movies := router.Group("/goteca")

	movies.GET("/movies", controller.Get())
	movies.POST("/movies", controller.Post())
	movies.DELETE("/movies/:id", controller.Delete())

	router.Run(":8080")
}