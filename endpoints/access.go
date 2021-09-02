package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/mes1234/golock/adapters"
	"github.com/mes1234/golock/service"
)

// Prepare endpoint for access service
func MakeEndpoint(svc service.AccessService, endpoint string) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		switch endpoint {
		case "addlocker":
			return svc.NewLocker(ctx, request.(adapters.AddLockerRequest))
		case "additem":
			return svc.Add(ctx, request.(adapters.AddItemRequest))
		case "removeitem":
			return svc.Remove(ctx, request.(adapters.RemoveItemRequest))
		case "getitem":
			return svc.Get(ctx, request.(adapters.GetItemRequest))
		default:
			panic("wrong endpoint name")
		}
	}
}
