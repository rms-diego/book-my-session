package authhandler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	authdto "github.com/rms-diego/book-my-session/internal/modules/auth/dto"
	authservice "github.com/rms-diego/book-my-session/internal/modules/auth/service"
	"github.com/rms-diego/book-my-session/internal/utils/validation"
	"github.com/rms-diego/book-my-session/pkg/config"
	"github.com/rms-diego/book-my-session/pkg/exception"
)

type authHandler struct {
	service authservice.AuthService
}

type AuthHandler interface {
	SignUp(c *gin.Context)
}

func NewAuthHandler(service authservice.AuthService) AuthHandler {
	return &authHandler{service}
}

func (h *authHandler) SignUp(c *gin.Context) {
	var payload authdto.SignUpRequest
	if err := validation.BindAndValidate(c, &payload); err != nil {
		c.Error(exception.NewException(err.Error(), http.StatusBadRequest))
		return
	}
	token, err := h.service.SignUp(payload)
	if err != nil {
		c.Error(err)
		return
	}

	exp := int(time.Now().Add(time.Hour * 12).Unix())
	c.SetCookie("Authorization", *token, exp, "/", config.Env.COOKIE_DOMAIN, false, true)
	c.JSON(http.StatusNoContent, nil)
}
