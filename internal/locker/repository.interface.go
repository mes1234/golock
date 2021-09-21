package locker

import "github.com/google/uuid"

type LockerRepository interface {
	// Create locker for given client
	InitLocker(lockerId uuid.UUID, resChan chan<- uuid.UUID)

	// Update locker in repository
	UpdateLocker(l Locker, lockerId uuid.UUID, resChan chan<- bool)

	// Retrieve locker
	GetLocker(lockerId uuid.UUID, resChan chan<- Locker)
}
