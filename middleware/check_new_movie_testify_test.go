package middleware

import (
	"testing"

	"github.com/ricardolonga/goteca"
	"github.com/ricardolonga/goteca/http"
	"github.com/stretchr/testify/assert"
)

func Test_InvalidMovie(t *testing.T) {
	movie := &goteca.Movie{Name: "Batman"}

	err := http.Validate(movie)
	assert.NotNil(t, err)
	assert.Equal(t, "Category is required.", err.Error())
}

func Test_ValidMovie(t *testing.T) {
	movie := &goteca.Movie{Name: "Batman", Category: "Action"}

	err := http.Validate(movie)
	assert.Nil(t, err)
}
