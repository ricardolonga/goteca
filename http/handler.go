package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricardolonga/goteca"
)

type handler struct {
	service goteca.Service
}

func NewHandler(service goteca.Service) http.Handler {
	handler := &handler{
		service: service,
	}

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(recovery())

	movies := router.Group("/goteca")
	movies.GET("/movies", count(), handler.getAll)
	movies.GET("/movies/:id", count(), handler.get)
	movies.POST("/movies", count(), CheckNewMovie, handler.Post)
	movies.DELETE("/movies/:id", count(), handler.delete)

	return router
}

func (me *handler) get(context *gin.Context) {
	id := context.Param("id")

	movie, err := me.service.Get(id)
	if err != nil {
		context.AbortWithStatus(http.StatusNotFound)
		return
	}

	context.JSON(http.StatusOK, movie)
}

func (me *handler) getAll(context *gin.Context) {
	movies, err := me.service.GetAll()
	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusOK, movies)
}

func (me *handler) Post(context *gin.Context) {
	movieInterface, ok := context.Get("movie")
	if !ok {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	savedMovie, err := me.service.Save(movieInterface.(*goteca.Movie))
	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusOK, savedMovie)
}

func (me *handler) delete(context *gin.Context) {
	id := context.Param("id")

	if err := me.service.Delete(id); err != nil {
		context.AbortWithStatus(http.StatusNotFound)
		return
	}

	context.AbortWithStatus(http.StatusOK)
}
