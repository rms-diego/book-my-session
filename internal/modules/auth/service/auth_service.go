package authservice

import (
	"context"
	"net/http"

	"github.com/rms-diego/book-my-session/internal/model"
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
	SignUp(ctx context.Context, dto authdto.SignUpRequest) (*string, error)
	SignIn(ctx context.Context, dto authdto.SignInRequest) (*string, error)
	RefreshToken(ctx context.Context, token string) (*string, error)
}

func NewAuthService(repository authrepository.AuthRepository) AuthService {
	return &authService{repository}
}

func (s *authService) SignUp(ctx context.Context, data authdto.SignUpRequest) (*string, error) {
	user, err := s.repository.GetByEmail(ctx, data.Email)
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

	uc, err := s.repository.Create(ctx, u)
	if err != nil {
		return nil, err
	}

	strToken, err := token.GenerateToken(*uc)
	if err != nil {
		return nil, err
	}

	return &strToken, nil
}

func (s *authService) SignIn(ctx context.Context, data authdto.SignInRequest) (*string, error) {
	user, err := s.repository.GetByEmail(ctx, data.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, exception.NewException("user not found", http.StatusNotFound)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		return nil, exception.NewException("invalid credentials", http.StatusUnauthorized)
	}

	uf := model.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}

	strToken, err := token.GenerateToken(uf)
	if err != nil {
		return nil, err
	}

	return &strToken, nil
}

func (s *authService) RefreshToken(ctx context.Context, strToken string) (*string, error) {
	claims, err := token.ValidateAndDecodeToken(strToken)
	if err != nil {
		return nil, err
	}

	newClaims := model.User{
		ID:    claims.ID,
		Name:  claims.Name,
		Email: claims.Email,
		Role:  claims.Role,
	}

	newToken, err := token.GenerateToken(newClaims)
	if err != nil {
		return nil, err
	}

	return &newToken, nil
}
