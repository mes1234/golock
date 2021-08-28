package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
	"github.com/mes1234/golock/dto"
	"github.com/mes1234/golock/internal/client"
	"github.com/mes1234/golock/service"
)

// Container for all exported endpoints
type Endpoints struct {
}

func makeAccessEndpoint(svc service.Access) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.AddLockerRequest)

		clientId, err := uuid.Parse(req.Client)

		if err != nil {
			return dto.AddLockerResponse{
				LockerId: "",
				Err:      err.Error()}, nil
		}

		credentials := client.Credentials{
			Identity: client.Identity{
				Id: clientId,
			},
			Password: client.Password{
				Value: req.Password},
		}

		v, err := svc.NewLocker(ctx, credentials)

		if err != nil {
			return dto.AddLockerResponse{
				LockerId: v.String(),
				Err:      err.Error()}, nil
		}
		return dto.AddLockerResponse{
			LockerId: v.String(),
			Err:      ""}, nil
	}
}
