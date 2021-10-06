package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/mes1234/golock/adapters"
	"github.com/mes1234/golock/service"
)

// MakeTokenEndpoint prepare endpoint for access service
func MakeTokenEndpoint(svc service.TokenService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return svc.GetToken(ctx, request.(adapters.TokenRequest))
	}
}
