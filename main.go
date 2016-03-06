package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ricardolonga/golang/controller"
"github.com/ricardolonga/golang/repository"
	"github.com/ricardolonga/golang/middleware"
	"net/http"
)

func main() {
	repository := repository.New()

	router := gin.New()

	router.GET("/goteca", func(context *gin.Context) {
		context.AbortWithStatus(http.StatusOK)
	})

	movies := router.Group("/goteca")
	movies.Use(middleware.Log())

	movies.GET("/movies", controller.Get(repository))
	movies.POST("/movies", controller.Post(repository))
	movies.DELETE("/movies/:id", controller.Delete(repository))

	router.Run(":8080")
}