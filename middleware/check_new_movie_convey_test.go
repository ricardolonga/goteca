package middleware

import (
	"testing"

	"github.com/ricardolonga/goteca"
	"github.com/ricardolonga/goteca/http"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMovieCRUD(t *testing.T) {
	Convey("DADO que o usuário deseja cadastrar um novo filme", t, func() {
		Convey("QUANDO ele não informar a categoria", func() {
			movie := &goteca.Movie{Name: "Batman"}

			Convey("ENTAO a mensagem de erro deve ser 'Category is required", func() {
				err := http.Validate(movie)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "Category is required.")
			})
		})

		Convey("QUANDO ele informar a categoria", func() {
			movie := &goteca.Movie{Name: "Batman", Category: "Action"}

			Convey("ENTAO o filme deve ser valido", func() {
				err := http.Validate(movie)
				So(err, ShouldBeNil)
			})
		})
	})
}
