package locker

import "github.com/google/uuid"

type Repository interface {
	// Init  create locker for given client
	Init(lockerId uuid.UUID, resChan chan<- uuid.UUID)

	// Update locker in repository
	Update(l Locker, lockerId uuid.UUID, resChan chan<- bool)

	// Get locker
	Get(lockerId uuid.UUID, resChan chan<- Locker)
}
