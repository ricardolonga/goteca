package middleware

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"github.com/ricardolonga/goteca/entity"
)

func TestMiddleware(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Middleware Suite")
}

var _ = Describe("Check middlewares", func() {
	var movie *entity.Movie

	BeforeEach(func() {
		movie = &entity.Movie{ Name: "Batman" }
	})

	Describe("DADO que o usuário deseja cadastrar um novo filme", func() {
		Context("QUANDO ele não informar a categoria", func() {
			It("ENTAO a mensagem de erro deve ser 'Category is required.'", func() {
				err := validate(movie)
				Expect(err.Error()).To(Equal("Category is required."))
			})
		})

		Context("QUANDO ele informar a categoria", func() {
			It("ENTAO o filme deve ser valido.", func() {
				movie.Category = "Action"
				err := validate(movie)
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("DADO que o usuário deseja excluir um filme", func() {})
})
