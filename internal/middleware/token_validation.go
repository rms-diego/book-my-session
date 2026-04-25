package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rms-diego/book-my-session/internal/utils/token"
	"github.com/rms-diego/book-my-session/pkg/exception"
)

func ValidationToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		t, err := c.Cookie("Authorization")
		if err != nil {
			c.Error(exception.NewException("Unauthorized", http.StatusUnauthorized))
			c.Abort()
			return
		}

		claims, err := token.DecodeToken(t)
		if err != nil {
			c.Error(exception.NewException("Unauthorized", http.StatusUnauthorized))
			c.Abort()
			return
		}

		ctx := token.NewContext(c.Request.Context(), claims)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
