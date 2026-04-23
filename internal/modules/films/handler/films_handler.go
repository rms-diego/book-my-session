package filmshandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	filmsdto "github.com/rms-diego/book-my-session/internal/modules/films/dto"
	filmsservice "github.com/rms-diego/book-my-session/internal/modules/films/service"
	"github.com/rms-diego/book-my-session/internal/utils/validation"
)

type filmsHandler struct {
	service filmsservice.FilmsService
}

type FilmsHandler interface {
	CreateFilm(c *gin.Context)
}

func NewFilmsHandler(service filmsservice.FilmsService) FilmsHandler {
	return &filmsHandler{service}
}

func (h *filmsHandler) CreateFilm(c *gin.Context) {
	var payload filmsdto.CreateFilmRequest
	if err := validation.BindAndValidate(c, &payload); err != nil {
		c.Error(err)
		return
	}

	if err := h.service.CreateFilm(payload); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
