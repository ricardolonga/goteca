package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/ricardolonga/goteca/repository"
	"gopkg.in/mgo.v2"
)

const MOVIES = "movies"

func Get(repository repository.Repository) gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")

		movie, err := repository.Find(MOVIES, id)
		if err != nil {
			if err.Error() == mgo.ErrNotFound.Error() {
				context.AbortWithStatus(http.StatusNotFound)
				return
			}

			context.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		context.JSON(http.StatusOK, movie)
	}
}

func GetAll(repository repository.Repository) gin.HandlerFunc {
	return func(context *gin.Context) {
		movies, err := repository.FindAll(MOVIES)
		if err != nil {
			context.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		context.JSON(http.StatusOK, movies)
	}
}

func Post(repository repository.Repository) gin.HandlerFunc {
	return func(context *gin.Context) {
		movie, ok := context.Get("movie")
		if !ok {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		savedMovie, err := repository.Save(MOVIES, movie)
		if err != nil {
			context.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		context.JSON(http.StatusOK, savedMovie)
	}
}

func Delete(repository repository.Repository) gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")

		if err := repository.Delete(MOVIES, id); err != nil {
			if err.Error() == mgo.ErrNotFound.Error() {
				context.AbortWithStatus(http.StatusNotFound)
				return
			}

			context.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		context.AbortWithStatus(http.StatusOK)
	}
}