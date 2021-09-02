package service

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/mes1234/golock/dto"
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
	reques dto.AddItemRequest,
) (dto.AddItemResponse, error) {
	return dto.AddItemResponse{Status: true}, nil
}

// Get item from locker
func (s accessService) Get(
	ctx context.Context,
	request dto.GetItemRequest,
) (dto.GetItemResponse, error) {
	return dto.GetItemResponse{
		Content: locker.PlainContent{
			Value: make([]byte, 0),
		}}, nil
}

// Remove item from locker
func (s accessService) Remove(
	ctx context.Context,
	request dto.RemoveItemRequest,
) (dto.RemoveItemResponse, error) {
	return dto.RemoveItemResponse{
		Status: true,
	}, nil
}

// Add new locker
func (s accessService) NewLocker(
	ctx context.Context,
	request dto.AddLockerRequest, // Identification of client
) (dto.AddLockerResponse, error) {

	logger := log.With(s.logger, "method", "Add")
	logger.Log("Successfully added locker {id}", request.Client)

	response := dto.AddLockerResponse{
		LockerId: request.Client,
	}
	return response, nil
}
