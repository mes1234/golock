package locker

import "github.com/mes1234/golock/internal/keys"

type LockerRepository interface {
	// Create locker for given client
	AddLocker(
		resChan chan<- LockerId,
	)

	// Update locker in repository
	UpdateLocker(
		locker Locker,
	)

	// Retrieve locker
	GetLocker(
		lockerId LockerId,
	) Locker
}

type LockerManager interface {
	// Add item to locker
	AddItem(
		lockerId LockerId,
		secretName SecretId,
		key keys.Value,
		content PlainContent,
		resChan chan<- error)

	// Remove item from locker
	RemoveItem(
		lockerId LockerId,
		secretName SecretId,
		resChan chan<- error)

	// Get item from locker
	GetItem(
		lockerId LockerId,
		key keys.Value,
		secretName SecretId,
		resChan chan<- struct {
			PlainContent
			error
		})
}
