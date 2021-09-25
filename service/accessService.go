package service

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/google/uuid"
	"github.com/mes1234/golock/adapters"
	"github.com/mes1234/golock/internal/keys"
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

	lockerCh := make(chan locker.Locker)
	errCh := make(chan error)

	repo := locker.GetRepository(reques.ClientId)
	go repo.GetLocker(reques.LockerId, lockerCh)

	l, ok := <-lockerCh

	if !ok {
		return adapters.AddItemResponse{Status: false}, nil
	}

	go l.AddItem(
		reques.SecretId,
		keys.Value{},
		reques.Content,
		errCh)

	err := <-errCh

	go repo.UpdateLocker(l, reques.LockerId, make(chan<- bool))

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

	lockerCh := make(chan locker.Locker)
	go locker.GetRepository(request.ClientId).GetLocker(request.LockerId, lockerCh)

	l, ok := <-lockerCh
	if !ok {
		return adapters.GetItemResponse{Content: make([]byte, 0)}, errors.New("error getting access to locker")
	}

	contentCh := make(chan []byte)
	go l.GetItem(keys.Value{}, request.SecretId, contentCh)

	c, ok := <-contentCh
	if !ok {
		return adapters.GetItemResponse{Content: make([]byte, 0)}, errors.New("error getting content")
	}

	return adapters.GetItemResponse{Content: c}, nil
}

// Remove item from locker
func (s accessService) Remove(ctx context.Context,
	request adapters.RemoveItemRequest,
) (adapters.RemoveItemResponse, error) {
	return adapters.RemoveItemResponse{
		Status: true,
	}, nil
}

// NewLocker Add new locker
func (s accessService) NewLocker(
	ctx context.Context,
	request adapters.AddLockerRequest, // Identification of client
) (adapters.AddLockerResponse, error) {

	lockerCh := make(chan uuid.UUID)

	go locker.GetRepository(request.ClientId).InitLocker(uuid.New(), lockerCh)

	// Await response
	response := adapters.AddLockerResponse{
		LockerId: <-lockerCh,
	}
	return response, nil
}
