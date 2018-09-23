package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
