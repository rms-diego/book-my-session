package authservice

import (
	"net/http"

	authdto "github.com/rms-diego/book-my-session/internal/modules/auth/dto"
	authrepository "github.com/rms-diego/book-my-session/internal/modules/auth/repository"
	"github.com/rms-diego/book-my-session/internal/utils/token"
	"github.com/rms-diego/book-my-session/pkg/exception"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	repository authrepository.AuthRepository
}

type AuthService interface {
	SignUp(dto authdto.SignUpRequest) (*string, error)
}

func NewAuthService(repository authrepository.AuthRepository) AuthService {
	return &authService{repository}
}

func (s *authService) SignUp(data authdto.SignUpRequest) (*string, error) {
	user, err := s.repository.FindByEmail(data.Email)
	if err != nil {
		return nil, err
	}

	if user != nil {
		return nil, exception.NewException("user already exist", http.StatusUnprocessableEntity)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u := authdto.SignUpRequest{
		Email:    data.Email,
		Password: string(hashedPassword),
		Role:     data.Role,
		Name:     data.Name,
	}

	uc, err := s.repository.Create(u)
	if err != nil {
		return nil, err
	}

	strToken, err := token.GenerateToken(*uc)
	if err != nil {
		return nil, err
	}

	return &strToken, nil
}
