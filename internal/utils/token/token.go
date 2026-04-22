package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rms-diego/book-my-session/internal/model"
	"github.com/rms-diego/book-my-session/pkg/config"
)

type UserClaims struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(payload model.User) (string, error) {
	claims := jwt.MapClaims{
		"id":    payload.ID,
		"name":  payload.Name,
		"email": payload.Email,
		"role":  payload.Role,
		"exp":   time.Now().Add(time.Hour * 12).Unix(),
	}

	t := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return t.SignedString([]byte(config.Env.JWT_SECRET))
}

func DecodeToken[T jwt.Claims](tokenString string, claims T) (T, error) {
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return []byte(config.Env.JWT_SECRET), nil
	})

	if err != nil {
		var zero T
		return zero, err
	}

	return token.Claims.(T), nil
}

func ValidateToken(tokenString string) bool {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return []byte(config.Env.JWT_SECRET), nil
	})
	return err == nil
}
