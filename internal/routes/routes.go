package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rms-diego/book-my-session/internal/modules/auth"
	"github.com/rms-diego/book-my-session/internal/modules/films"
)

func Init(r *gin.Engine) {
	// general routes
	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "server is running"})
	})

	// auth routes
	authModule := auth.NewAuthModule()
	authModule.InitRoutes(r.Group("/auth"))

	// films routes
	filmsModule := films.NewFilmsModule()
	filmsModule.InitRoutes(r.Group("/films"))
}
