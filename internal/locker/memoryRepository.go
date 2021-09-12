package locker

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/mes1234/golock/internal/client"
	"github.com/mes1234/golock/internal/key"
)

type repository struct {
	mu          sync.Mutex
	repository  map[LockerId]Locker
	persistance *repository
}

var memoryRepository repository

func init() {
	memoryRepository = repository{
		repository:  make(map[LockerId]Locker),
		persistance: nil,
	}
}

func GetRepository() LockerRepository {
	return &memoryRepository
}

func (r *repository) AddLocker(
	clientId client.ClientId,
	resChan chan<- LockerId) {
	// Ensure thread safety
	r.mu.Lock()
	defer r.mu.Unlock()
	//
	newLocker := Locker{
		Id:      uuid.New(),
		Client:  clientId,
		Secrets: map[SecretId]Secret{},
	}
	r.repository[newLocker.Id] = newLocker
	//
	// Persist change
	//	go r.persistance.AddLocker(clientId)
	// return
	resChan <- newLocker.Id
}

func (r *repository) AddItem(
	clientId client.ClientId,
	lockerId LockerId,
	secretName SecretId,
	content PlainContent,
	resChan chan<- error) {
	// Ensure thread safety
	r.mu.Lock()
	defer r.mu.Unlock()
	//

	if _, ok := r.repository[lockerId]; !ok {
		resChan <- errors.New("this locker does not exists")
		return
	}
	if r.repository[lockerId].Client != clientId {
		resChan <- errors.New("this locker does not belong to You")
		return
	}

	var revision int16
	if current, ok := r.repository[lockerId].Secrets[secretName]; ok {
		revision = current.Revision
	} else {
		revision = 0
	}
	secret := NewCrypter().encrypt(clientId, key.Value{}, content)
	secret.Revision = revision
	r.repository[lockerId].Secrets[secretName] = secret

	//
	// Persist change
	//	go r.persistance.AddItem(clientId, lockerId, secretName, content)
	// return
	resChan <- nil
}

func (r *repository) RemoveItem(
	clientId client.ClientId,
	lockerId LockerId,
	secretName SecretId,
	resChan chan<- error) {
	// Ensure thread safety
	r.mu.Lock()
	defer r.mu.Unlock()
	//
	// DO LOGIC
	//
	// Persist change
	//	go r.persistance.RemoveItem(clientId, lockerId, secretName)
	// return
	resChan <- nil
}

func (r *repository) GetItem(
	clientId client.ClientId,
	lockerId LockerId,
	secretName SecretId,
	resChan chan<- struct {
		PlainContent
		error
	}) {
	// Ensure thread safety
	r.mu.Lock()
	defer r.mu.Unlock()
	//
	// DO LOGIC
	//
	// Persist change
	// Do this if local stage failed r.persistance.GetItem(clientId, lockerId, secretName)
	// return
	resChan <- struct {
		PlainContent
		error
	}{PlainContent{}, nil}
}
