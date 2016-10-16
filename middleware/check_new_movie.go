package middleware

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
	"github.com/ricardolonga/goteca/entity"
)

func CheckNewMovie() gin.HandlerFunc {
	return func(context *gin.Context) {
		movie := &entity.Movie{}

		if err := context.BindJSON(movie); err != nil {
			context.AbortWithError(http.StatusBadRequest, err)
			return
		}

		if err := validate(movie); err != nil {
			context.AbortWithError(http.StatusBadRequest, err)
			return
		}

		context.Set("movie", movie)
		context.Next()
	}
}

func validate(movie *entity.Movie) error {
	if movie.Category == "" {
		return errors.New("Category is required.")
	}

	return nil
}
