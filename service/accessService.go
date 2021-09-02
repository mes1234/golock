package service

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/mes1234/golock/dto"
	"github.com/mes1234/golock/internal/client"
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
	client client.Credentials, // Identification of client
	lockerId locker.LockerId, // Identification of locker to insert into
	secretId locker.SecretId, // Identification of secret to get
) (locker.PlainContent, error) {
	return locker.PlainContent{
		Value: make([]byte, 0),
	}, nil
}

// Delete item from locker
func (s accessService) Deleted(
	ctx context.Context,
	client client.Credentials, // Identification of client
	lockerId locker.LockerId, // Identification of locker to insert into
	secretId locker.SecretId, // Identification of secret to get
) (bool, error) {
	return true, nil
}

// Add new locker
func (s accessService) NewLocker(
	ctx context.Context,
	request dto.AddLockerRequest, // Identification of client
) (dto.AddLockerResponse, error) {

	logger := log.With(s.logger, "method", "Add")
	logger.Log("Successfully added locker {id}", request.Client.Identity.Id)

	response := dto.AddLockerResponse{
		LockerId: request.Client.Identity.Id,
	}
	return response, nil
}
