package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/NeowayLabs/logger"
)

func Get() gin.HandlerFunc {
	return func(context *gin.Context) {
		logger.Info("Getting all movies...")
		context.AbortWithStatus(http.StatusOK)
	}
}

func Post() gin.HandlerFunc {
	return func(context *gin.Context) {
		logger.Info("Creating a movie...")
		context.AbortWithStatus(http.StatusOK)
	}
}

func Delete() gin.HandlerFunc {
	return func(context *gin.Context) {
		logger.Info("Deleting a movie...")
		context.AbortWithStatus(http.StatusOK)
	}
}