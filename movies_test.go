package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostMovies(t *testing.T) {
	server := httptest.NewServer(newHandler(newMongoSession()))

	/*
	 * Post
	 */
	movieBytes := []byte(`{
		"name": "Batman"
	}`)

	post, err := http.NewRequest(http.MethodPost, server.URL+"/goteca/movies", bytes.NewReader(movieBytes))
	assert.NoError(t, err)

	res, err := http.DefaultClient.Do(post)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)

	resultBytes, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, resultBytes)

	var result map[string]interface{}
	err = json.Unmarshal(resultBytes, &result)
	assert.NoError(t, err)

	movieID, _ := result["id"].(string)
	assert.NotEmpty(t, movieID)

	/*
	 * Get
	 */

	get, err := http.NewRequest(http.MethodGet, server.URL+"/goteca/movies/"+movieID, nil)

	res, err = http.DefaultClient.Do(get)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	resultBytes, err = ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, resultBytes)

	var getResult map[string]interface{}
	err = json.Unmarshal(resultBytes, &getResult)
	assert.NoError(t, err)

	movieName, _ := getResult["name"].(string)
	assert.Equal(t, "Batman", movieName)

}
