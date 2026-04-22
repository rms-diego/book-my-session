package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "server is running"})
	})
}
