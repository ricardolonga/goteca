package middleware

import (
	"testing"

	"github.com/ricardolonga/goteca"
	"github.com/ricardolonga/goteca/http"
)

func Test_InvalidMovie_Native(t *testing.T) {
	movie := &goteca.Movie{Name: "Batman"}

	err := http.Validate(movie)

	if err == nil || err.Error() != "Category is required." {
		t.Fail()
	}
}
