package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ricardolonga/goteca/controller"
	"github.com/ricardolonga/goteca/middleware"
	"github.com/ricardolonga/goteca/repository"
	"gopkg.in/mgo.v2"
)

func main() {
	session, _ := mgo.Dial(os.Getenv("MONGO_URL"))

	repository := repository.New(session)

	router := gin.New()

	movies := router.Group("/goteca")
	movies.GET("/movies", controller.GetAll(repository))
	movies.GET("/movies/:id", controller.Get(repository))
	movies.POST("/movies", middleware.CheckNewMovie(), controller.Post(repository))
	movies.DELETE("/movies/:id", controller.Delete(repository))

	router.Run()
}
