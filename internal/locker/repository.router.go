package locker

import (
	"github.com/google/uuid"
	"github.com/mes1234/golock/internal/client"
)

type repositoryRouter struct {
	repositories []LockerRepository
}

func GetRepository(clientId client.ClientId) LockerRepository {
	repo := repositoryRouter{}
	repo.repositories = append(repo.repositories, getMemoryRepository(clientId))
	repo.repositories = append(repo.repositories, getDbRepository(clientId))

	return &repo
}

func (r *repositoryRouter) UpdateLocker(locker Locker, lockerId uuid.UUID, resChan chan<- bool) {
	for _, v := range r.repositories {
		go v.UpdateLocker(locker, lockerId, make(chan<- bool))
	}
	resChan <- true

}

func (r *repositoryRouter) GetLocker(lockerId uuid.UUID, resChan chan<- Locker) {

	for _, v := range r.repositories {
		lockerCh := make(chan Locker)
		go v.GetLocker(lockerId, lockerCh)
		res, ok := <-lockerCh
		if ok {
			res.IncreaseRevision()
			resChan <- res
			break
		}
	}
	close(resChan)
}

func (r *repositoryRouter) InitLocker(lockerId uuid.UUID, resChan chan<- uuid.UUID) {
	for _, v := range r.repositories {
		lockerCh := make(chan uuid.UUID)
		go v.InitLocker(lockerId, lockerCh)
		res, ok := <-lockerCh
		if !ok {
			resChan <- res
			break
		}

	}
	resChan <- lockerId
}
