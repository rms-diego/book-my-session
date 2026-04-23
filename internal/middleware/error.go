package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rms-diego/book-my-session/pkg/exception"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err
		message := err.Error()
		if customErr, ok := err.(exception.Exception); ok {
			c.AbortWithStatusJSON(customErr.Code(), gin.H{"error": message})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": message})
	}
}
