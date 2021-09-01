package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/mes1234/golock/dto"
	"github.com/mes1234/golock/service"
)

// Prepare endpoint for access service
func MakeEndpoint(svc service.AccessService, endpoint string) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		switch endpoint {
		case "addlocker":
			v := request.(dto.AddLockerRequest)
			return handleAddLockerRequest(ctx, svc, v)
		default:
			panic("wrong endpoint name")
		}
	}
}

// Handler for AddLockerRequest
func handleAddLockerRequest(
	ctx context.Context,
	svc service.AccessService,
	request dto.AddLockerRequest) (interface{}, error) {
	v, err := svc.NewLocker(ctx, request.Client)

	// handle error
	if err != nil {
		return nil, err
	}

	// return response
	return dto.AddLockerResponse{
		LockerId: v,
	}, nil
}
