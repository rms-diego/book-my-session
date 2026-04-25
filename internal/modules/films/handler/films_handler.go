package filmshandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	filmsdto "github.com/rms-diego/book-my-session/internal/modules/films/dto"
	filmsservice "github.com/rms-diego/book-my-session/internal/modules/films/service"
	"github.com/rms-diego/book-my-session/internal/shared"
	"github.com/rms-diego/book-my-session/internal/utils/validation"
)

type filmsHandler struct {
	service filmsservice.FilmsService
}

type FilmsHandler interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
}

func NewFilmsHandler(service filmsservice.FilmsService) FilmsHandler {
	return &filmsHandler{service}
}

func (h *filmsHandler) Create(c *gin.Context) {
	var payload filmsdto.CreateFilmRequest
	if err := validation.BindAndValidateBody(c, &payload); err != nil {
		c.Error(err)
		return
	}

	if err := h.service.Create(payload); err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *filmsHandler) Update(c *gin.Context) {
	var payload filmsdto.UpdateFilmRequest
	var params shared.IDParam
	if err := validation.BindAndValidateBody(c, &payload); err != nil {
		c.Error(err)
		return
	}

	if err := validation.BindAndValidateParams(c, &params); err != nil {
		c.Error(err)
		return
	}

	if err := h.service.Update(params.ID, payload); err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}
