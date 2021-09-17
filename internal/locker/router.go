package locker

import "github.com/mes1234/golock/internal/client"

type lockerRouter struct {
	repositories []LockerRepository
}

func GetRepository(clientId client.ClientId) LockerRepository {
	repo := lockerRouter{}
	repo.repositories = append(repo.repositories, getMemoryRepository(clientId))
	repo.repositories = append(repo.repositories, getDbRepository(clientId))

	return &repo
}

func (r *lockerRouter) UpdateLocker(locker Locker) {
	for _, v := range r.repositories {
		go v.UpdateLocker(locker)
	}

}

func (r *lockerRouter) GetLocker(lockerId LockerId, resChan chan<- Locker) {

	for _, v := range r.repositories {
		lockerCh := make(chan Locker)
		go v.GetLocker(lockerId, lockerCh)
		res, ok := <-lockerCh
		if ok {
			resChan <- res
			break
		}
	}
}

func (r *lockerRouter) InitLocker(lockerId LockerId, resChan chan<- LockerId) {
	for _, v := range r.repositories {
		lockerCh := make(chan LockerId)
		go v.InitLocker(lockerId, lockerCh)
		res, ok := <-lockerCh
		if !ok {
			resChan <- res
			break
		}

	}
	resChan <- lockerId
}
