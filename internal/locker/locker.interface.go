package locker

import (
	"github.com/mes1234/golock/internal/keys"
)

type Locker interface {
	GetId() LockerId
	IncreaseRevision()

	// Add item to locker
	AddItem(
		secretName SecretId,
		key keys.Value,
		content PlainContent,
		resChan chan<- error)

	// Remove item from locker
	RemoveItem(
		secretName SecretId,
		resChan chan<- error)

	// Get item from locker
	GetItem(
		key keys.Value,
		secretName SecretId,
		resChan chan<- PlainContent)
}
