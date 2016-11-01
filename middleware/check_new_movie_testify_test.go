package middleware

import (
    "testing"
    "github.com/ricardolonga/goteca/entity"
    "github.com/stretchr/testify/assert"
)

func Test_InvalidMovie(t *testing.T) {
    movie := &entity.Movie{ Name: "Batman" }

    err := validate(movie)
    assert.NotNil(t, err)
    assert.Equal(t, "Category is required.", err.Error())
}

func Test_ValidMovie(t *testing.T) {
    movie := &entity.Movie{ Name: "Batman", Category: "Action" }

    err := validate(movie)
    assert.Nil(t, err)
}