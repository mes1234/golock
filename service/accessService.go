package service

import (
	"context"

	"github.com/go-kit/kit/log"
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
	client client.Credentials, // Identification of client
	lockerId locker.LockerId, // Identification of locker to insert into
	secretId locker.SecretId, // Identification of secret to get
	content locker.PlainContent, // Content which shall be injected
) (bool, error) {
	return true, nil
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
	client client.Credentials, // Identification of client
) (locker.LockerId, error) {

	logger := log.With(s.logger, "method", "Add")
	logger.Log("Successfully added locker {id}", client.Identity.Id)

	// time.Sleep(100 * time.Millisecond)
	return client.Identity.Id, nil
}
