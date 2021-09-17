package locker

import (
	"sync"

	"github.com/google/uuid"
	"github.com/mes1234/golock/internal/client"
)

type memoryRepository struct {
	r  map[LockerId]Locker
	mu *sync.Mutex
	c  client.ClientId
}

var memRepository map[LockerId]Locker

func init() {
	memRepository = make(map[LockerId]Locker)
}

func getMemoryRepository(clientId client.ClientId) LockerRepository {

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

func (r *memoryRepository) InitLocker(lockerId LockerId, resChan chan<- LockerId) {

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

	resChan <- newLocker.Id
}
