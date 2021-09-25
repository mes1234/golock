package locker

import (
	"github.com/google/uuid"
	"github.com/mes1234/golock/internal/client"
)

type repositoryRouter struct {
	repositories []Repository
}

func GetRepository(clientId client.Id) Repository {
	repo := repositoryRouter{}
	repo.repositories = append(repo.repositories, getMemoryRepository(clientId))
	repo.repositories = append(repo.repositories, getDbRepository(clientId))

	return &repo
}

func (r *repositoryRouter) Update(locker Locker, lockerId uuid.UUID, resChan chan<- bool) {
	for _, v := range r.repositories {
		go v.Update(locker, lockerId, make(chan<- bool))
	}
	resChan <- true

}

func (r *repositoryRouter) Get(lockerId uuid.UUID, resChan chan<- Locker) {

	for _, v := range r.repositories {
		lockerCh := make(chan Locker)
		go v.Get(lockerId, lockerCh)
		res, ok := <-lockerCh
		if ok {
			res.IncreaseRevision()
			resChan <- res
			break
		}
	}
	close(resChan)
}

func (r *repositoryRouter) Init(lockerId uuid.UUID, resChan chan<- uuid.UUID) {
	for _, v := range r.repositories {
		lockerCh := make(chan uuid.UUID)
		go v.Init(lockerId, lockerCh)
		res, ok := <-lockerCh
		if !ok {
			resChan <- res
			break
		}

	}
	resChan <- lockerId
}
