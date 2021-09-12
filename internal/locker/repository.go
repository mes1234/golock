package locker

import "github.com/mes1234/golock/internal/client"

type LockerRepository interface {
	// Create locker for given client
	AddLocker(
		clientId client.ClientId,
		resChan chan<- LockerId,
	)

	// Add item to locker
	AddItem(
		clientId client.ClientId,
		lockerId LockerId,
		secretName SecretId,
		content PlainContent,
		resChan chan<- error)

	// Remove item from locker
	RemoveItem(
		clientId client.ClientId,
		lockerId LockerId,
		secretName SecretId,
		resChan chan<- error)

	// Get item from locker
	GetItem(
		clientId client.ClientId,
		lockerId LockerId,
		secretName SecretId,
		resChan chan<- struct {
			PlainContent
			error
		})
}
