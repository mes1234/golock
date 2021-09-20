package locker

import (
	"sync"

	"github.com/google/uuid"
	"github.com/mes1234/golock/internal/client"
)

type memoryRepository struct {
	r  map[uuid.UUID]Locker
	mu *sync.Mutex
	c  client.ClientId
}

var memRepository map[uuid.UUID]Locker

func init() {
	memRepository = make(map[uuid.UUID]Locker)
}

func getMemoryRepository(clientId client.ClientId) LockerRepository {

	return &memoryRepository{
		r:  memRepository,
		mu: &sync.Mutex{},
		c:  clientId,
	}
}

func (r *memoryRepository) UpdateLocker(locker Locker, lockerId uuid.UUID, resChan chan<- bool) {
	r.r[lockerId] = locker
	resChan <- true
}

func (r *memoryRepository) GetLocker(lockerId uuid.UUID, resChan chan<- Locker) {
	// Ensure thread safety
	r.mu.Lock()
	defer r.mu.Unlock()

	if l, ok := r.r[lockerId]; ok {
		resChan <- l
	} else {
		close(resChan)
	}
}

func (r *memoryRepository) InitLocker(lockerId uuid.UUID, resChan chan<- uuid.UUID) {

	// Ensure thread safety
	r.mu.Lock()
	defer r.mu.Unlock()

	newLocker := GetMemoryLocker(r.c, lockerId)

	r.r[lockerId] = newLocker

	resChan <- lockerId
}
