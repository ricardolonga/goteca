package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gitlab.com/ricardolonga/goteca/repository"
	"github.com/NeowayLabs/logger"
)

func Get(repository repository.Repository) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.AbortWithStatus(http.StatusOK)
	}
}

func Post(repository repository.Repository) gin.HandlerFunc {
	return func(context *gin.Context) {
		logger.Info("POST ok...")
		context.AbortWithStatus(http.StatusOK)
	}
}

func Delete(repository repository.Repository) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.AbortWithStatus(http.StatusOK)
	}
}