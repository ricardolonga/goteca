package middleware

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/ricardolonga/goteca/entity"
	"github.com/NeowayLabs/logger"
	"net/http"
)

func CheckNewMovie() gin.HandlerFunc {
	return func(context *gin.Context) {
		movie := &entity.Movie{}

		if err := context.BindJSON(movie); err != nil {
			logger.Error("Error on unmarshal a new movie: %s", err)
			context.AbortWithError(http.StatusBadRequest, err)
			return
		}

		if movie.Category == "" {
			logger.Error("Invalid new movie, category is undefined.")
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		context.Set("movie", movie)

		context.Next()
	}
}
