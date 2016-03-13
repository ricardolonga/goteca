package main

import (
	"testing"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"net/http"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"github.com/ricardolonga/goteca/controller"
	"strings"
	"github.com/ricardolonga/goteca/middleware"
)

type InvalidMovieMockRepository struct {
	T *testing.T
}

func (me *InvalidMovieMockRepository) Save(collection string, object interface{}) (savedObject interface{}, err error) {
	assert.Fail(me.T, "Nao deveria ter chamado este metodo...")
	return
}

func (me *InvalidMovieMockRepository) FindAll(collection string) (objects []interface{}, err error) {
	assert.Fail(me.T, "Nao deveria ter chamado este metodo...")
	return
}

func (me *InvalidMovieMockRepository) Find(collection string, id string) (object interface{}, err error) {
	assert.Fail(me.T, "Nao deveria ter chamado este metodo...")
	return
}

func (me *InvalidMovieMockRepository) Delete(collection string, id string) (err error) {
	assert.Fail(me.T, "Nao deveria ter chamado este metodo...")
	return
}

func Test_invalid_new_movie(t *testing.T) {
	router := gin.New()

	movies := router.Group("/goteca")
	movies.POST("/movies", middleware.CheckNewMovie(), controller.Post(&InvalidMovieMockRepository{T: t}))

	w := httptest.NewRecorder()
	invalidMovieBytes := readJsonFile(t, "datasets/invalid_movie.json")
	req, _ := http.NewRequest("POST", "/goteca/movies", strings.NewReader(string(invalidMovieBytes)))
	router.ServeHTTP(w, req)

	body, err := ioutil.ReadAll(w.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Empty(t, body)
}