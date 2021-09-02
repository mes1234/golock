package service

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/mes1234/golock/adapters"
	"github.com/mes1234/golock/internal/locker"
)

type accessService struct {
	logger log.Logger
}

func NewAccessService(log log.Logger) AccessService {
	return &accessService{
		logger: log,
	}
}

// Add item to locker
func (s accessService) Add(
	ctx context.Context,
	reques adapters.AddItemRequest,
) (adapters.AddItemResponse, error) {
	return adapters.AddItemResponse{Status: true}, nil
}

// Get item from locker
func (s accessService) Get(
	ctx context.Context,
	request adapters.GetItemRequest,
) (adapters.GetItemResponse, error) {
	return adapters.GetItemResponse{
		Content: locker.PlainContent{
			Value: []byte("hello"),
		}}, nil
}

// Remove item from locker
func (s accessService) Remove(
	ctx context.Context,
	request adapters.RemoveItemRequest,
) (adapters.RemoveItemResponse, error) {
	return adapters.RemoveItemResponse{
		Status: true,
	}, nil
}

// Add new locker
func (s accessService) NewLocker(
	ctx context.Context,
	request adapters.AddLockerRequest, // Identification of client
) (adapters.AddLockerResponse, error) {

	logger := log.With(s.logger, "method", "Add")
	logger.Log("Successfully added locker {id}", request.ClientId)

	response := adapters.AddLockerResponse{
		LockerId: request.ClientId,
	}
	return response, nil
}
