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

	lockerCh := make(chan error)

	go locker.GetRepository().AddItem(reques.ClientId, reques.LockerId, reques.SecretId, reques.Content, lockerCh)

	err := <-lockerCh
	var status bool
	if err != nil {
		status = false
	} else {
		status = true
	}

	return adapters.AddItemResponse{Status: status}, nil
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

	lockerCh := make(chan locker.LockerId)

	go locker.GetRepository().AddLocker(request.ClientId, lockerCh)

	// Await response
	response := adapters.AddLockerResponse{
		LockerId: <-lockerCh,
	}
	return response, nil
}
