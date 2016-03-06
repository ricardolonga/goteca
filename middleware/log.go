package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/NeowayLabs/logger"
	"time"
)

func Log() gin.HandlerFunc {
	return func(context *gin.Context) {
		now := time.Now()
		context.Next()
		logger.Info("Method: %s - Uri: %s - Time: %s", context.Request.Method, context.Request.RequestURI, time.Since(now))
	}
}
