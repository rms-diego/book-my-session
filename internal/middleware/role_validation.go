package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rms-diego/book-my-session/internal/utils/token"
	"github.com/rms-diego/book-my-session/pkg/exception"
)

func ValidateRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		claims, exists := token.FromContext(ctx)
		if !exists {
			c.Error(exception.NewException("Unauthorized", http.StatusUnauthorized))
			c.Abort()
			return
		}

		if claims.Role != "admin" {
			c.Error(exception.NewException("Forbidden", http.StatusForbidden))
			c.Abort()
			return
		}

		c.Next()
	}
}
