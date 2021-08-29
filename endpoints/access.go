package endpoints

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/mes1234/golock/dto"
	"github.com/mes1234/golock/service"
)

// Prepare endpoint for access service
func MakeAccessEndpoint(svc service.Access) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		switch v := request.(type) {
		case dto.AddLockerRequest:
			return handleAddLockerRequest(ctx, svc, v)
		default:
			return dto.AddLockerResponse{},
				errors.New("not supported request")
		}
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
		return dto.AddLockerResponse{
			Err: err,
		}, err
	}

	// return response
	return dto.AddLockerResponse{
		LockerId: v,
	}, nil
}
