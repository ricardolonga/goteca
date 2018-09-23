package middleware

import (
	"testing"

	"github.com/ricardolonga/goteca"
	"github.com/ricardolonga/goteca/http"
	"github.com/stretchr/testify/suite"
)

type CheckMiddlewareSuite struct {
	suite.Suite
	movie *goteca.Movie
}

func (suite *CheckMiddlewareSuite) SetupSuite() {}

func (suite *CheckMiddlewareSuite) SetupTest() {
	suite.movie = &goteca.Movie{Name: "Batman"}
}

func (suite *CheckMiddlewareSuite) TestInvalidMovie() {
	err := http.Validate(suite.movie)
	suite.NotNil(err)
	suite.Equal("Category is required.", err.Error())
}

func (suite *CheckMiddlewareSuite) TearDownTest() {}

func (suite *CheckMiddlewareSuite) TearDownSuite() {}

func TestCheckMiddlewareSuite(t *testing.T) {
	suite.Run(t, new(CheckMiddlewareSuite))
}
