package service

import (
	"context"

	"github.com/mes1234/golock/internal/client"
	"github.com/mes1234/golock/internal/locker"
)

//Access interface describes a service that insert and retrieve data
type AccessService interface {

	// Add item to locker
	Add(
		ctx context.Context,
		client client.Credentials, // Identification of client
		lockerId locker.LockerId, // Identification of locker to insert into
		secretId locker.SecretId, // Identification of secret to get
		content locker.PlainContent, // Content which shall be injected
	) (bool, error) // status of operation

	// Get item from locker
	Get(
		ctx context.Context,
		client client.Credentials, // Identification of client
		lockerId locker.LockerId, // Identification of locker to insert into
		secretId locker.SecretId, // Identification of secret to get
	) (locker.PlainContent, error) // uncrypted content

	// Delete item from locker
	Deleted(
		ctx context.Context,
		client client.Credentials, // Identification of client
		lockerId locker.LockerId, // Identification of locker to insert into
		secretId locker.SecretId, // Identification of secret to get
	) (bool, error) // status of operation

	// Add new locker
	NewLocker(
		ctx context.Context,
		client client.Credentials, // Identification of client
	) (locker.LockerId, error) // Identification of locker
}
