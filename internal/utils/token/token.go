package token

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rms-diego/book-my-session/internal/model"
	"github.com/rms-diego/book-my-session/pkg/config"
)

type UserClaims struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

type claims string

const CONTEXT_CLAIMS_KEY claims = "claims"

func GenerateToken(payload model.User) (string, error) {
	t := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		UserClaims{
			ID:    payload.ID,
			Name:  payload.Name,
			Email: payload.Email,
			Role:  payload.Role,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 12)),
			},
		},
	)

	return t.SignedString([]byte(config.Env.JWT_SECRET))
}

func DecodeToken(tokenString string) (*UserClaims, error) {
	claims := &UserClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return []byte(config.Env.JWT_SECRET), nil
	})

	if err != nil {
		return nil, err
	}

	userClaims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, jwt.ErrTokenMalformed
	}

	return userClaims, nil
}

func ValidateToken(tokenString string) bool {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return []byte(config.Env.JWT_SECRET), nil
	})
	return err == nil
}

func FromContext(ctx context.Context) (*UserClaims, bool) {
	claims, ok := ctx.Value(CONTEXT_CLAIMS_KEY).(*UserClaims)
	return claims, ok
}

func NewContext(ctx context.Context, claims *UserClaims) context.Context {
	return context.WithValue(ctx, CONTEXT_CLAIMS_KEY, claims)
}
