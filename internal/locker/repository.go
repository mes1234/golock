package locker

type LockerRepository interface {
	// Create locker for given client
	InitLocker(lockerId LockerId, resChan chan<- LockerId)

	// Update locker in repository
	UpdateLocker(locker Locker)

	// Retrieve locker
	GetLocker(LockerId, chan<- Locker)
}
