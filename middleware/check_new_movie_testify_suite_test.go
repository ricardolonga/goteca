package middleware

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/ricardolonga/goteca/entity"
)

type CheckMiddlewareSuite struct {
	suite.Suite
	movie *entity.Movie
}

func (suite *CheckMiddlewareSuite) SetupSuite() {}

func (suite *CheckMiddlewareSuite) SetupTest() {
	suite.movie = &entity.Movie{ Name: "Batman" }
}

func (suite *CheckMiddlewareSuite) TestInvalidMovie() {
	err := validate(suite.movie)
	suite.NotNil(err)
	suite.Equal("Category is required.", err.Error())
}

func (suite *CheckMiddlewareSuite) TearDownTest() {}

func (suite *CheckMiddlewareSuite) TearDownSuite() {}

func TestCheckMiddlewareSuite(t *testing.T) {
	suite.Run(t, new(CheckMiddlewareSuite))
}
