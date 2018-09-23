package http

import (
	"io/ioutil"
	netHttp "net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ricardolonga/goteca"
	"github.com/stretchr/testify/assert"
)

type InvalidMovieMockRepository struct {
	T *testing.T
}

func (me *InvalidMovieMockRepository) Save(object *goteca.Movie) (*goteca.Movie, error) {
	assert.Fail(me.T, "Nao deveria ter chamado este metodo...")
	return nil, nil
}

func (me *InvalidMovieMockRepository) FindAll() ([]*goteca.Movie, error) {
	assert.Fail(me.T, "Nao deveria ter chamado este metodo...")
	return nil, nil
}

func (me *InvalidMovieMockRepository) Find(id string) (*goteca.Movie, error) {
	assert.Fail(me.T, "Nao deveria ter chamado este metodo...")
	return nil, nil
}

func (me *InvalidMovieMockRepository) Delete(id string) error {
	assert.Fail(me.T, "Nao deveria ter chamado este metodo...")
	return nil
}

func Test_invalid_new_movie(t *testing.T) {
	service := goteca.NewService(&InvalidMovieMockRepository{T: t})
	handler := NewHandler(service)

	w := httptest.NewRecorder()
	invalidMovieBytes := readJsonFile(t, "../datasets/invalid_movie.json")
	req, _ := netHttp.NewRequest("POST", "/goteca/movies", strings.NewReader(string(invalidMovieBytes)))
	handler.ServeHTTP(w, req)

	body, err := ioutil.ReadAll(w.Body)
	assert.Nil(t, err)
	assert.Equal(t, netHttp.StatusBadRequest, w.Code)
	assert.Empty(t, body)
}
