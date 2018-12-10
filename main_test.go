package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"gopkg.in/mgo.v2"
)

var session *mgo.Session

func init() {
	var err error

	session, err = mgo.Dial("172.20.0.2:27017")
	if err != nil {
		return
	}
}

func TestMoviesWithRecorder(t *testing.T) {
	router := newHandler(session)

	req, _ := http.NewRequest(http.MethodPost, "/goteca/movies", bytes.NewReader([]byte(`{
		"name": "Batman"
	}`)))

	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	assert.Equal(t, http.StatusCreated, res.Code)

	resBytes, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)

	var movie map[string]interface{}
	assert.NoError(t, json.Unmarshal(resBytes, &movie))
	assert.NotEmpty(t, movie["_id"])

	movieID := movie["_id"].(string)

	get, _ := http.NewRequest(http.MethodGet, "/goteca/movies/"+movieID, nil)

	res = httptest.NewRecorder()
	router.ServeHTTP(res, get)
	assert.Equal(t, http.StatusOK, res.Code)

	resBytes, err = ioutil.ReadAll(res.Body)
	assert.NoError(t, err)

	var getMovie map[string]interface{}
	assert.NoError(t, json.Unmarshal(resBytes, &getMovie))

	assert.NotEmpty(t, movie["_id"])
	assert.Equal(t, movie["_id"], getMovie["_id"])
}

func TestMoviesWithServer(t *testing.T) {
	server := httptest.NewServer(newHandler(session))
	defer server.Close()

	req, _ := http.NewRequest(http.MethodPost, server.URL+"/goteca/movies", bytes.NewReader([]byte(`{
		"name": "Batman"
	}`)))

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)

	resBytes, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)

	var movie map[string]interface{}
	assert.NoError(t, json.Unmarshal(resBytes, &movie))
	assert.NotEmpty(t, movie["_id"])

	movieID := movie["_id"].(string)

	get, _ := http.NewRequest(http.MethodGet, server.URL+"/goteca/movies/"+movieID, nil)

	res, err = http.DefaultClient.Do(get)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	resBytes, err = ioutil.ReadAll(res.Body)
	assert.NoError(t, err)

	var getMovie map[string]interface{}
	assert.NoError(t, json.Unmarshal(resBytes, &getMovie))

	assert.NotEmpty(t, movie["_id"])
	assert.Equal(t, movie["_id"], getMovie["_id"])
}

func TestMovies(t *testing.T) {
	t.Run("not found", func(t *testing.T) {
		router := newHandler(session)

		get := httptest.NewRequest(http.MethodGet, "/goteca/movies/1", nil)

		res := httptest.NewRecorder()
		router.ServeHTTP(res, get)
		assert.Equal(t, http.StatusNotFound, res.Code)

		resBytes, err := ioutil.ReadAll(res.Body)
		assert.NoError(t, err)
		assert.Empty(t, resBytes)
	})

	t.Run("invalid json post", func(t *testing.T) {
		server := httptest.NewServer(newHandler(session))
		defer server.Close()

		req, _ := http.NewRequest(http.MethodPost, server.URL+"/goteca/movies", bytes.NewReader([]byte(`{
			...
		}`)))

		res, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)

		resBytes, err := ioutil.ReadAll(res.Body)
		assert.NoError(t, err)
		assert.Empty(t, resBytes)
	})
}
