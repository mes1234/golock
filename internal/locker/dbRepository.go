package locker

import "github.com/mes1234/golock/internal/client"

type dbRepository struct {
}

func getDbRepository(clientId client.ClientId) LockerRepository {

	return &dbRepository{}
}

func (r *dbRepository) UpdateLocker(locker Locker) {

}

func (r *dbRepository) GetLocker(lockerId LockerId, resChan chan<- Locker) {
	close(resChan)
}

func (r *dbRepository) InitLocker(lockerId LockerId, resChan chan<- LockerId) {
	resChan <- lockerId
}
