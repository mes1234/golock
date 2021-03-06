package locker

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mes1234/golock/internal/client"
	"github.com/mes1234/golock/internal/keys"
)

// Locker is container for all secrets
type memoryLocker struct {
	Crypter  Crypter           // provide cryptographic functionality
	Revision int               // revision of current locker
	Id       uuid.UUID         // Identifier of locker
	Client   uuid.UUID         // Identifiers of all clients with access
	Secrets  map[string]Secret //Content of Locker

}

func GetMemoryLocker(clientId client.Id, lockerId uuid.UUID) Locker {
	return &memoryLocker{
		Id:       lockerId,
		Revision: 1,
		Client:   clientId,
		Secrets:  map[string]Secret{},
		Crypter:  NewCrypter(),
	}
}

func (r *memoryLocker) ItemsToCommit() map[string]Secret {
	items := make(map[string]Secret)
	for k, v := range r.Secrets {
		if v.Revision == r.Revision {
			items[k] = v
		}
	}
	return items
}

func (r *memoryLocker) GetId() uuid.UUID {
	return r.Id
}

func (r *memoryLocker) GetClientId() uuid.UUID {
	return r.Client
}

func (r *memoryLocker) IncreaseRevision() {
	r.Revision = r.Revision + 1
}

// AddItem adds item to locker
func (r *memoryLocker) AddItem(
	secretName string,
	key keys.Value,
	content []byte,
	revision int,
	resChan chan<- error) {

	secret := r.Crypter.encrypt(key, content)
	secret.Revision = r.Revision
	secret.Active = true

	// Nonzero revision forces locker to adjust to new one
	if revision != 0 {
		secret.Revision = revision
		r.Revision = revision
	}
	r.Secrets[secretName] = secret

	resChan <- nil
}

// RemoveItem removes item from locker
func (r *memoryLocker) RemoveItem(
	secretName string,
	resChan chan<- error) {
	//
	if _, ok := r.Secrets[secretName]; !ok {
		resChan <- errors.New("no item found for given secret id")
	}
	if !r.Secrets[secretName].Active {
		resChan <- errors.New("item already removed")
	}
	item := r.Secrets[secretName]
	item.Active = false
	r.Secrets[secretName] = item

	resChan <- nil
}

// GetItem gets item from locker
func (r *memoryLocker) GetItem(
	key keys.Value,
	secretName string,
	resChan chan<- []byte) {
	// Ensure thread safety
	//
	if _, ok := r.Secrets[secretName]; !ok {
		close(resChan)
		return
	}
	if !r.Secrets[secretName].Active {
		close(resChan)
		return
	}
	resChan <- r.Crypter.decrypt(key, r.Secrets[secretName])
}
