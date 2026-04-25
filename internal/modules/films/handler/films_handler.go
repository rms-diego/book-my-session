package filmshandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	filmsdto "github.com/rms-diego/book-my-session/internal/modules/films/dto"
	filmsservice "github.com/rms-diego/book-my-session/internal/modules/films/service"
	"github.com/rms-diego/book-my-session/internal/shared"
	"github.com/rms-diego/book-my-session/internal/utils/validation"
	"github.com/rms-diego/book-my-session/pkg/exception"
)

type filmsHandler struct {
	service filmsservice.FilmsService
}

type FilmsHandler interface {
	Create(c *gin.Context)
	UploadThumbnail(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
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

	if err := h.service.Create(c.Request.Context(), payload); err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *filmsHandler) Update(c *gin.Context) {
	var params shared.IDParam
	if err := validation.BindAndValidateParams(c, &params); err != nil {
		c.Error(err)
		return
	}

	var payload filmsdto.UpdateFilmRequest
	if err := validation.BindAndValidateBody(c, &payload); err != nil {
		c.Error(err)
		return
	}

	if err := h.service.Update(c.Request.Context(), params.ID, payload); err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *filmsHandler) Delete(c *gin.Context) {
	var params shared.IDParam
	if err := validation.BindAndValidateParams(c, &params); err != nil {
		c.Error(err)
		return
	}

	if err := h.service.Delete(c.Request.Context(), params.ID); err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *filmsHandler) GetAll(c *gin.Context) {
	films, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, films)
}

func (h *filmsHandler) GetById(c *gin.Context) {
	var params shared.IDParam
	if err := validation.BindAndValidateParams(c, &params); err != nil {
		c.Error(err)
		return
	}

	film, err := h.service.GetById(c.Request.Context(), params.ID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, film)
}

func (h *filmsHandler) UploadThumbnail(c *gin.Context) {
	var params shared.IDParam
	if err := validation.BindAndValidateParams(c, &params); err != nil {
		c.Error(err)
		return
	}

	f, err := c.FormFile("thumbnail")
	if err != nil {
		c.Error(exception.NewException("thumbnail file is required", http.StatusBadRequest))
		return
	}

	if err := h.service.UploadThumbnail(c.Request.Context(), params.ID, f); err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusOK)
}
