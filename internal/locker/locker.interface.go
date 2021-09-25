package locker

import (
	"github.com/google/uuid"
	"github.com/mes1234/golock/internal/keys"
)

type Locker interface {
	IncreaseRevision()
	GetClientId() uuid.UUID
	ItemsToCommit() map[string]Secret

	// AddItem Add item to locker
	AddItem(
		secretName string,
		key keys.Value,
		content []byte,
		revision int,
		resChan chan<- error)

	// RemoveItem Remove item from locker
	RemoveItem(
		secretName string,
		resChan chan<- error)

	// GetItem Get item from locker
	GetItem(
		key keys.Value,
		secretName string,
		resChan chan<- []byte)
}
