package middleware

import (
    "testing"
    "github.com/ricardolonga/goteca/entity"
)

func Test_InvalidMovie_Native(t *testing.T) {
    movie := &entity.Movie{ Name: "Batman" }

    err := validate(movie)

    if err == nil || err.Error() != "Category is required." {
        t.Fail()
    }
}