package middlewares

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/google/uuid"
	"github.com/mes1234/golock/adapters"
	"github.com/mes1234/golock/persistance"

	"github.com/dgrijalva/jwt-go"
	gokitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
)

// Timing middleware decorator
func AuthorizationMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {

			data := ctx.Value(gokitjwt.JWTClaimsContextKey).(*jwt.StandardClaims)
			genericRequest := request.(adapters.ClientAssigner)

			clientId, _ := uuid.Parse(data.Id)

			repository := persistance.NewClientRepository()

			client := adapters.Client{
				ClientId: clientId,
			}

			repository.Retrieve(&client)

			request = genericRequest.AssignClient(client.ClientName)

			logger := log.With(logger, "method", "Add")
			logger.Log("Successfully authorized user {user}", client.ClientName)

			// request = genericRequest.(interface{})

			return next(ctx, request)
		}
	}
}
