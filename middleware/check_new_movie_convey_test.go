package middleware

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/ricardolonga/goteca/entity"
)

func TestMovieCRUD(t *testing.T) {
	Convey("DADO que o usuário deseja cadastrar um novo filme", t, func() {
		Convey("QUANDO ele não informar a categoria", func() {
			movie := &entity.Movie{ Name: "Batman" }

			Convey("ENTAO a mensagem de erro deve ser 'Category is required", func() {
				err := validate(movie)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "Category is required.")
			})
		})

		Convey("QUANDO ele informar a categoria", func() {
			movie := &entity.Movie{ Name: "Batman", Category: "Action" }

			Convey("ENTAO o filme deve ser valido", func() {
				err := validate(movie)
				So(err, ShouldBeNil)
			})
		})
	})
}
