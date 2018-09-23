package http

import (
	"net/http/httptest"
	"testing"

	netHttp "net/http"

	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/ricardolonga/goteca"
	"github.com/stretchr/testify/assert"
)

type NewMovieMockRepository struct {
	T *testing.T
}

func (me *NewMovieMockRepository) Save(movie *goteca.Movie) (*goteca.Movie, error) {
	movie.Id = "1"
	return movie, nil
}

func (me *NewMovieMockRepository) FindAll() ([]*goteca.Movie, error) {
	assert.Fail(me.T, "Nao deveria ter chamado este metodo...")
	return nil, nil
}

func (me *NewMovieMockRepository) Find(id string) (*goteca.Movie, error) {
	assert.Fail(me.T, "Nao deveria ter chamado este metodo...")
	return nil, nil
}

func (me *NewMovieMockRepository) Delete(id string) error {
	assert.Fail(me.T, "Nao deveria ter chamado este metodo...")
	return nil
}

func Test_new_movie(t *testing.T) {
	service := goteca.NewService(&NewMovieMockRepository{T: t})
	handler := NewHandler(service)

	w := httptest.NewRecorder()
	invalidMovieBytes := []byte("{ \"name\": \"Man of Fire\", \"category\": \"Action\" }")
	req, _ := netHttp.NewRequest("POST", "/goteca/movies", bytes.NewReader(invalidMovieBytes))
	handler.ServeHTTP(w, req)

	body, err := ioutil.ReadAll(w.Body)
	assert.Nil(t, err)
	assert.Equal(t, netHttp.StatusOK, w.Code)
	assert.NotNil(t, body)

	savedMovie := &goteca.Movie{}
	json.Unmarshal(body, &savedMovie)

	assert.NotNil(t, savedMovie)

	assert.Equal(t, "1", savedMovie.Id)
	assert.Equal(t, "Action", savedMovie.Category)
	assert.Equal(t, "Man of Fire", savedMovie.Name)
}
