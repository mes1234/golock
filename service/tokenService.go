package service

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/mes1234/golock/adapters"
	"github.com/mes1234/golock/persistance"
)

type tokenService struct {
	logger log.Logger
}

func (t tokenService) GetToken(
	ctx context.Context,
	request adapters.TokenRequest,
) (adapters.TokenResponse, error) {

	repository := persistance.NewClientRepository()

	client := adapters.Client{
		ClientName: request.Username,
	}

	repository.Retrieve(&client)

	return adapters.TokenResponse{
		Token: getToken(client.ClientId),
	}, nil
}

func NewTokenService(log log.Logger) TokenService {
	return &tokenService{
		logger: log,
	}
}

func getToken(clientId uuid.UUID) string {
	key := []byte("test")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.StandardClaims{
			Id: clientId.String(),
		})
	tokenString, _ := token.SignedString(key)
	return tokenString
}
