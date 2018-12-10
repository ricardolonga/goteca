package main

import (
	"fmt"
	"net/http"
	"os"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"

	"github.com/gin-gonic/gin"
)

func getMovies(session *mgo.Session) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		ID := c.Param("id")

		s := session.Copy()
		defer s.Close()

		var movie map[string]interface{}
		if err := session.DB("goteca").C("movies").FindId(ID).One(&movie); err != nil {
			if err == mgo.ErrNotFound {
				c.AbortWithStatus(http.StatusNotFound)
				return
			}

			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, movie)
	}
}

func postMovies(session *mgo.Session) func(c *gin.Context) {
	return func(c *gin.Context) {
		var movie map[string]interface{}
		if err := c.BindJSON(&movie); err != nil {
			return
		}

		s := session.Copy()
		defer s.Close()

		movie["_id"] = bson.NewObjectId().Hex()
		session.DB("goteca").C("movies").Insert(movie)

		c.JSON(http.StatusCreated, movie)
	}
}

func newHandler(session *mgo.Session) *gin.Engine {
	router := gin.New()

	goteca := router.Group("/goteca")
	goteca.GET("/movies/:id", getMovies(session))
	goteca.POST("/movies", postMovies(session))

	return router
}

func main() {
	mongoURL := os.Getenv("MONGO_URL")
	if mongoURL == "" {
		fmt.Printf("environment variable MONGO_URL is required")
		return
	}

	session, err := mgo.Dial(mongoURL)
	if err != nil {
		fmt.Printf("error on MongoDB connection: %q", err)
		return
	}

	newHandler(session).Run()
}
