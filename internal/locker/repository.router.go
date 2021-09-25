package locker

import (
	"github.com/google/uuid"
	"github.com/mes1234/golock/internal/client"
)

type repositoryRouter struct {
	memory Repository
	db     Repository
}

func GetRepository(clientId client.Id) Repository {
	repo := repositoryRouter{}
	repo.memory = getMemoryRepository(clientId)
	repo.db = getDbRepository(clientId)

	return &repo
}

func (r *repositoryRouter) Update(locker Locker, lockerId uuid.UUID, resChan chan<- bool) {

	go r.memory.Update(locker, lockerId, make(chan<- bool))
	go r.db.Update(locker, lockerId, make(chan<- bool))

	resChan <- true

}

func (r *repositoryRouter) getFromMemory(lockerId uuid.UUID, resChan chan<- Locker) {
	// Get Data from Memory
	lockerCh := make(chan Locker)
	go r.memory.Get(lockerId, lockerCh)
	res, ok := <-lockerCh
	if ok {
		res.IncreaseRevision()
		resChan <- res
		return
	}

	// If no Db provider then end processing
	if r.db == nil {
		close(resChan)
		return
	}

	// Recover locker from Db
	go r.getFromDb(lockerId, r.memory.Update, resChan)

}

func (r *repositoryRouter) getFromDb(lockerId uuid.UUID, updateMemory func(l Locker, lockerId uuid.UUID, resChan chan<- bool), resChan chan<- Locker) {
	lockerCh := make(chan Locker)
	go r.db.Get(lockerId, lockerCh)
	res, ok := <-lockerCh
	if ok {
		go updateMemory(res, lockerId, make(chan bool, 0))
		resChan <- res
	} else {
		close(resChan)
	}

}

func (r *repositoryRouter) Get(lockerId uuid.UUID, resChan chan<- Locker) {

	go r.getFromMemory(lockerId, resChan)
}

func (r *repositoryRouter) Create(lockerId uuid.UUID, resChan chan<- uuid.UUID) {

	lockerCh := make(chan uuid.UUID)
	go r.memory.Create(lockerId, lockerCh)
	res, ok := <-lockerCh
	if !ok {
		resChan <- res
	}
	resChan <- lockerId
}
