package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gitlab.com/ricardolonga/goteca/repository"
)

func Get(repository repository.Repository) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.AbortWithStatus(http.StatusOK)
	}
}

func Post(repository repository.Repository) gin.HandlerFunc {
	return func(context *gin.Context) {
		movie, ok := context.Get("movie")
		if !ok {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		savedMovie, err := repository.Save("movies", movie)
		if err != nil {
			context.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		context.JSON(http.StatusOK, savedMovie)
	}
}

func Delete(repository repository.Repository) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.AbortWithStatus(http.StatusOK)
	}
}