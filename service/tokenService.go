package service

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/golang-jwt/jwt"
	"github.com/mes1234/golock/adapters"
)

type tokenService struct {
	logger log.Logger
}

func (t tokenService) GetToken(
	ctx context.Context,
	request adapters.TokenRequest,
) (adapters.TokenResponse, error) {
	return adapters.TokenResponse{
		Token: getToken(),
	}, nil
}

func NewTokenService(log log.Logger) TokenService {
	return &tokenService{
		logger: log,
	}
}

func getToken() string {
	key := []byte("test")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.StandardClaims{
			Id: "b48f5b98-e9e7-40ce-b8cf-cdc4d2c59061",
		})
	tokenString, _ := token.SignedString(key)
	return tokenString
}
