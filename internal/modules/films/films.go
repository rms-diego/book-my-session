package films

import (
	"github.com/gin-gonic/gin"
	"github.com/rms-diego/book-my-session/internal/database"
	"github.com/rms-diego/book-my-session/internal/middleware"
	filmshandler "github.com/rms-diego/book-my-session/internal/modules/films/handler"
	filmsrepository "github.com/rms-diego/book-my-session/internal/modules/films/repository"
	filmsservice "github.com/rms-diego/book-my-session/internal/modules/films/service"
)

type filmsModule struct {
	handler    filmshandler.FilmsHandler
	Repository filmsrepository.FilmsRepository
	Service    filmsservice.FilmsService
}

type FilmsModule interface {
	InitRoutes(r *gin.RouterGroup)
}

func NewFilmsModule() FilmsModule {
	r := filmsrepository.NewFilmsRepository(database.Db)
	s := filmsservice.NewFilmsService(r)
	h := filmshandler.NewFilmsHandler(s)

	return &filmsModule{
		handler:    h,
		Repository: r,
		Service:    s,
	}
}

func (m *filmsModule) InitRoutes(r *gin.RouterGroup) {
	r.Use(middleware.ValidationToken())

	r.POST("/", middleware.ValidateRole(), m.handler.Create)
	r.PUT("/:id", middleware.ValidateRole(), m.handler.Update)
}
