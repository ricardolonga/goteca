package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ricardolonga/goteca"
	"net/http"
)

func CheckNewMovie(context *gin.Context) {
	movie := &goteca.Movie{}

	if err := context.BindJSON(movie); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := Validate(movie); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	context.Set("movie", movie)
	context.Next()
}

func Validate(movie *goteca.Movie) error {
	if movie.Category == "" {
		return errors.New("Category is required.")
	}

	return nil
}
