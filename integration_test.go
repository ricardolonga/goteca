package main

import (
	"testing"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"net/http"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"encoding/json"
	"github.com/NeowayLabs/logger"
	"gitlab.com/ricardolonga/goteca/controller"
	"strings"
)

type MockRepository struct {
	T *testing.T
}

func (me *MockRepository) Save() {
	assert.Fail(me.T, "Nao deveria ter chamado este metodo...")
}

func (me *MockRepository) Find() {
	assert.Fail(me.T, "Nao deveria ter chamado este metodo...")
}

func (me *MockRepository) Delete() {
	assert.Fail(me.T, "Nao deveria ter chamado este metodo...")
}

func Test_invalid_new_movie(t *testing.T) {
	router := gin.New()

	movies := router.Group("/goteca")
	movies.POST("/movies", controller.Post(&MockRepository{T: t}))

	w := httptest.NewRecorder()
	invalidMovieBytes := readJsonFile(t, "datasets/invalid_movie.json")
	req, _ := http.NewRequest("POST", "/goteca/movies", strings.NewReader(string(invalidMovieBytes)))
	router.ServeHTTP(w, req)

	body, err := ioutil.ReadAll(w.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, body)

	result := make(map[string]interface{}, 0)
	json.Unmarshal(body, &result)
	assert.NotNil(t, result)

	expected := make(map[string]interface{}, 0)
	expectedBytes, err := ioutil.ReadFile("expects/invalid_movie_expect.json")
	assert.Nil(t, err)

	json.Unmarshal(expectedBytes, &expected)

	if !assert.Equal(t, expected, result) {
		expectedJson, _ := json.Marshal(expected)
		logger.Info("Experado: %s", string(expectedJson))

		resultJson, _ := json.Marshal(result)
		logger.Info("Resultado: %s", string(resultJson))

		t.Fail()
	}
}

func readJsonFile(t *testing.T, filePath string) []byte {
	objectBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Error on read file [%s]. %s", filePath, err)
	}

	return objectBytes
}