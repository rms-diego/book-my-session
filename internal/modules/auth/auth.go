package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/rms-diego/book-my-session/internal/database"
	authhandler "github.com/rms-diego/book-my-session/internal/modules/auth/handler"
	authrepository "github.com/rms-diego/book-my-session/internal/modules/auth/repository"
	authservice "github.com/rms-diego/book-my-session/internal/modules/auth/service"
)

type authModule struct {
	handler    authhandler.AuthHandler
	Service    authservice.AuthService
	Repository authrepository.AuthRepository
}

type AuthModule interface {
	InitRoutes(r *gin.RouterGroup)
}

func NewAuthModule() AuthModule {
	r := authrepository.NewAuthRepository(database.Db)
	s := authservice.NewAuthService(r)
	h := authhandler.NewAuthHandler(s)

	return &authModule{
		handler:    h,
		Service:    s,
		Repository: r,
	}
}

func (m *authModule) InitRoutes(r *gin.RouterGroup) {
	r.POST("/sign-up", m.handler.SignUp)
}
