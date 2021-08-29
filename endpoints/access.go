package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/mes1234/golock/dto"
	"github.com/mes1234/golock/service"
)

// Prepare endpoint for access service
func MakeAddLockerEndpoint(svc service.Access) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		v := request.(dto.AddLockerRequest)
		return handleAddLockerRequest(ctx, svc, v)

	}
}

// Handler for AddLockerRequest
func handleAddLockerRequest(
	ctx context.Context,
	svc service.Access,
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
