package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rms-diego/book-my-session/pkg/exception"
	"go.uber.org/zap"
)

func ErrorHandler(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err

		if customErr, ok := err.(exception.Exception); ok {
			if customErr.Code() >= 400 && customErr.Code() < 500 {
				c.JSON(customErr.Code(), gin.H{
					"error": customErr.Error(),
				})

				return
			}
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}
