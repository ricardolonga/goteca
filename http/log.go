package http

import (
	"github.com/NeowayLabs/logger"
	"github.com/gin-gonic/gin"
	"time"
)

func count() gin.HandlerFunc {
	return func(context *gin.Context) {
		now := time.Now()
		context.Next()
		logger.Info("Method: %s - Uri: %s - Time: %s", context.Request.Method, context.Request.RequestURI, time.Since(now))
	}
}
