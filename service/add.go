package service

import (
	"context"

	"github.com/mes1234/golock/internal/client"
	"github.com/mes1234/golock/internal/locker"
)

// Add item to locker
func (AccessService) Add(
	ctx context.Context,
	client client.Credentials, // Identification of client
	lockerId locker.LockerId, // Identification of locker to insert into
	secretId locker.SecretId, // Identification of secret to get
	content locker.PlainContent, // Content which shall be injected
) (bool, error) {
	return true, nil
}

// Get item from locker
func (AccessService) Get(
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
func (AccessService) Deleted(
	ctx context.Context,
	client client.Credentials, // Identification of client
	lockerId locker.LockerId, // Identification of locker to insert into
	secretId locker.SecretId, // Identification of secret to get
) (bool, error) {
	return true, nil
}

// Add new locker
func (AccessService) NewLocker(
	ctx context.Context,
	client client.Credentials, // Identification of client
) (locker.LockerId, error) {
	return client.Identity.Id, nil
}
