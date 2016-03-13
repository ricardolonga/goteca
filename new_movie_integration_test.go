package main

import (
	"testing"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"net/http"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"encoding/json"
	"github.com/ricardolonga/goteca/controller"
	"strings"
	"github.com/ricardolonga/goteca/middleware"
	"github.com/ricardolonga/goteca/entity"
)

type NewMovieMockRepository struct {
	T *testing.T
}

func (me *NewMovieMockRepository) Save(collection string, object interface{}) (savedObject interface{}, err error) {
	assert.Equal(me.T, "movies", collection)
	movie := object.(*entity.Movie)
	movie.Id = "1"
	return object, nil
}

func (me *NewMovieMockRepository) FindAll(collection string) (objects []interface{}, err error) {
	assert.Fail(me.T, "Nao deveria ter chamado este metodo...")
	return
}

func (me *NewMovieMockRepository) Find(collection string, id string) (object interface{}, err error) {
	assert.Fail(me.T, "Nao deveria ter chamado este metodo...")
	return
}

func (me *NewMovieMockRepository) Delete(collection string, id string) (err error) {
	assert.Fail(me.T, "Nao deveria ter chamado este metodo...")
	return
}

func Test_new_movie(t *testing.T) {
	router := gin.New()

	movies := router.Group("/goteca")
	movies.POST("/movies", middleware.CheckNewMovie(), controller.Post(&NewMovieMockRepository{T: t}))

	w := httptest.NewRecorder()
	invalidMovieBytes := readJsonFile(t, "datasets/valid_movie.json")
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