package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rms-diego/book-my-session/internal/utils/token"
	"github.com/rms-diego/book-my-session/pkg/exception"
)

func ValidateRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqToken, err := c.Cookie("Authorization")
		if err != nil {
			c.Error(exception.NewException("Unauthorized", http.StatusUnauthorized))
			c.Abort()
			return
		}

		decodedToken, err := token.DecodeToken(reqToken)
		if err != nil {
			c.Error(exception.NewException("Unauthorized", http.StatusUnauthorized))
			c.Abort()
			return
		}

		if decodedToken.Role != "admin" {
			c.Error(exception.NewException("Forbidden", http.StatusForbidden))
			c.Abort()
			return
		}

		c.Next()
	}
}
