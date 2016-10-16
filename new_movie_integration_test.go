package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ricardolonga/goteca/controller"
	"github.com/ricardolonga/goteca/entity"
	"github.com/ricardolonga/goteca/middleware"
	"github.com/stretchr/testify/assert"
)

type NewMovieMockRepository struct {
	T *testing.T
}

func (me *NewMovieMockRepository) Save(collection string, object interface{}) (interface{}, error) {
	assert.Equal(me.T, "movies", collection)
	movie := object.(*entity.Movie)
	movie.Id = "1"
	return object, nil
}

func (me *NewMovieMockRepository) FindAll(collection string) ([]interface{}, error) {
	assert.Fail(me.T, "Nao deveria ter chamado este metodo...")
	return nil, nil
}

func (me *NewMovieMockRepository) Find(collection string, id string) (interface{}, error) {
	assert.Fail(me.T, "Nao deveria ter chamado este metodo...")
	return nil, nil
}

func (me *NewMovieMockRepository) Delete(collection string, id string) error {
	assert.Fail(me.T, "Nao deveria ter chamado este metodo...")
	return nil
}

func Test_new_movie(t *testing.T) {
	router := gin.New()

	movies := router.Group("/goteca")
	movies.POST("/movies", middleware.CheckNewMovie(), controller.Post(&NewMovieMockRepository{T: t}))

	w := httptest.NewRecorder()
	invalidMovieBytes := []byte("{ \"name\": \"Man of Fire\", \"category\": \"Action\" }")
	req, _ := http.NewRequest("POST", "/goteca/movies", strings.NewReader(string(invalidMovieBytes)))
	router.ServeHTTP(w, req)

	body, err := ioutil.ReadAll(w.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, body)

	savedMovie := &entity.Movie{}
	json.Unmarshal(body, &savedMovie)

	assert.NotNil(t, savedMovie)

	assert.Equal(t, "1", savedMovie.Id)
	assert.Equal(t, "Action", savedMovie.Category)
	assert.Equal(t, "Man of Fire", savedMovie.Name)
}
