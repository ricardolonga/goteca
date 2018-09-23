package middleware

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ricardolonga/goteca"
	"github.com/ricardolonga/goteca/http"
)

func TestMiddleware(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Middleware Suite")
}

var _ = Describe("Check middlewares", func() {
	var movie *goteca.Movie

	BeforeEach(func() {
		movie = &goteca.Movie{Name: "Batman"}
	})

	Describe("DADO que o usuário deseja cadastrar um novo filme", func() {
		Context("QUANDO ele não informar a categoria", func() {
			It("ENTAO a mensagem de erro deve ser 'Category is required.'", func() {
				err := http.Validate(movie)
				Expect(err.Error()).To(Equal("Category is required."))
			})
		})

		Context("QUANDO ele informar a categoria", func() {
			It("ENTAO o filme deve ser valido.", func() {
				movie.Category = "Action"
				err := http.Validate(movie)
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("DADO que o usuário deseja excluir um filme", func() {})
})
