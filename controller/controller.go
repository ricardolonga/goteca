package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/ricardolonga/golang/repository"
)

func Get(repository repository.Repository) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.AbortWithStatus(http.StatusOK)
	}
}

func Post(repository repository.Repository) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.AbortWithStatus(http.StatusOK)
	}
}

func Delete(repository repository.Repository) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.AbortWithStatus(http.StatusOK)
	}
}