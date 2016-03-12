package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/ricardolonga/goteca/controller"
	"gitlab.com/ricardolonga/goteca/repository"
	"gitlab.com/ricardolonga/goteca/middleware"
	"net/http"
"gopkg.in/mgo.v2"
	"os"
	"github.com/NeowayLabs/logger"
	"time"
)

func main() {
	repository := repository.New(GetSession())

	router := gin.New()

	router.GET("/goteca", func(context *gin.Context) {
		context.AbortWithStatus(http.StatusOK)
	})

	movies := router.Group("/goteca")
	movies.Use(middleware.Log())

	movies.GET("/movies", controller.Get(repository))
	movies.POST("/movies", middleware.CheckNewMovie(), controller.Post(repository))
	movies.DELETE("/movies/:id", controller.Delete(repository))

	router.Run(":8080")
}

func GetSession() *mgo.Session {
	var session *mgo.Session
	var err error

	session, err = mgo.Dial(os.Getenv("MONGO_URL"))

	if err != nil {
		logger.Error("Error on getting Mongo session! %s", err)

		for err != nil {
			time.Sleep(time.Minute)
			session, err = mgo.Dial(os.Getenv("MONGO_URL"))
		}
	}

	logger.Info("Conectado no Mongo com sucesso!")

	session.SetMode(mgo.Monotonic, true)

	return session
}