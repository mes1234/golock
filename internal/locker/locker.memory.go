package locker

import (
	"errors"

	"github.com/mes1234/golock/internal/client"
	"github.com/mes1234/golock/internal/keys"
)

// Locker is container for all secrect
type memoryLocker struct {
	Crypter  Crypter             // provide cryptgraphic functionality
	Revision int                 // revision of current locker
	Id       LockerId            // Identifier of locker
	Client   client.ClientId     // Identifiers of all clients with access
	Secrets  map[SecretId]Secret //Content of Locker

}

func GetMemoryLocker(clientId client.ClientId, lockerId LockerId) Locker {
	return &memoryLocker{
		Id:       lockerId,
		Revision: 1,
		Client:   clientId,
		Secrets:  map[SecretId]Secret{},
		Crypter:  NewCrypter(),
	}
}

func (r *memoryLocker) GetId() LockerId {
	return r.Id
}

func (r *memoryLocker) IncreaseRevision() {
	r.Revision = r.Revision + 1
}

// Add item to locker
func (r *memoryLocker) AddItem(
	secretName SecretId,
	key keys.Value,
	content PlainContent,
	resChan chan<- error) {

	secret := r.Crypter.encrypt(keys.Value{}, content)
	secret.Revision = r.Revision
	secret.Active = true
	r.Secrets[secretName] = secret

	resChan <- nil
}

// Remove item from locker
func (r *memoryLocker) RemoveItem(
	secretName SecretId,
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

// Get item from locker
func (r *memoryLocker) GetItem(
	key keys.Value,
	secretName SecretId,
	resChan chan<- PlainContent) {
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
