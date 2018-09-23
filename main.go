package main

import (
	"github.com/NeowayLabs/logger"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"os"
)

func newHandler(session *mgo.Session) *gin.Engine {
	router := gin.Default()

	movies := router.Group("/goteca")

	movies.GET("/movies", func(c *gin.Context) {
		s := session.Copy()
		defer s.Close()

		var movies []map[string]interface{}
		if err := s.DB("goteca").C("movies").Find(nil).All(&movies); err != nil {
			if mgo.ErrNotFound == err {
				c.AbortWithError(http.StatusNotFound, err)
				return
			}

			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, movies)
	})
	movies.GET("/movies/:id", func(c *gin.Context) {
		s := session.Copy()
		defer s.Close()

		var movie map[string]interface{}
		if err := s.DB("goteca").C("movies").FindId(c.Param("id")).One(&movie); err != nil {
			if mgo.ErrNotFound == err {
				c.AbortWithError(http.StatusNotFound, err)
				return
			}

			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, movie)
	})
	movies.POST("/movies", func(c *gin.Context) {
		s := session.Copy()
		defer s.Close()

		var movie map[string]interface{}
		if err := c.BindJSON(&movie); err != nil {
			return
		}

		movie["_id"] = bson.NewObjectId().Hex()

		if err := s.DB("goteca").C("movies").Insert(movie); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"id": movie["_id"],
		})
	})

	return router
}

func newMongoSession() *mgo.Session {
	session, err := mgo.Dial(os.Getenv("MONGO_URL"))
	if err != nil {
		logger.Fatal("erro ao conectar com o mongo: %q", err)
	}

	return session
}

func main() {
	newHandler(newMongoSession()).Run()
}
