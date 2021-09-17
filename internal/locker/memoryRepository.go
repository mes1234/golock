package locker

import (
	"sync"

	"github.com/google/uuid"
	"github.com/mes1234/golock/internal/client"
)

type LockerRepository interface {
	// Create locker for given client
	InitLocker(resChan chan<- LockerId)

	// Update locker in repository
	UpdateLocker(locker Locker)

	// Retrieve locker
	GetLocker(LockerId, chan<- Locker)
}

type memoryRepository struct {
	r  map[LockerId]Locker
	mu *sync.Mutex
	c  client.ClientId
}

var memRepository map[LockerId]Locker

func init() {
	memRepository = make(map[LockerId]Locker)
}

func GetRepository(clientId client.ClientId) LockerRepository {

	return &memoryRepository{
		r:  memRepository,
		mu: &sync.Mutex{},
		c:  clientId,
	}
}

func (r *memoryRepository) UpdateLocker(locker Locker) {
	r.r[locker.Id] = locker
}

func (r *memoryRepository) GetLocker(lockerId LockerId, resChan chan<- Locker) {
	// Ensure thread safety
	r.mu.Lock()
	defer r.mu.Unlock()

	if l, ok := r.r[lockerId]; ok {
		resChan <- l
	} else {
		close(resChan)
	}

}

func (r *memoryRepository) InitLocker(resChan chan<- LockerId) {

	// Ensure thread safety
	r.mu.Lock()
	defer r.mu.Unlock()

	newLocker := Locker{
		Id:      uuid.New(),
		Client:  r.c,
		Secrets: map[SecretId]Secret{},
		Crypter: NewCrypter(),
	}

	r.r[newLocker.Id] = newLocker
	//
	// Persist change
	//	go r.persistance.AddLocker(clientId)
	// return
	resChan <- newLocker.Id
}
