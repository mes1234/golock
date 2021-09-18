package locker

import (
	"errors"

	"github.com/google/uuid"
	"github.com/mes1234/golock/internal/client"
	"github.com/mes1234/golock/internal/keys"
)

// Locker is container for all secrect
type memoryLocker struct {
	Crypter Crypter             // provide cryptgraphic functionality
	Id      LockerId            // Identifier of locker
	Client  client.ClientId     // Identifiers of all clients with access
	Secrets map[SecretId]Secret //Content of Locker

}

func GetMemoryLocker(clientId client.ClientId) Locker {
	return &memoryLocker{
		Id:      uuid.New(),
		Client:  clientId,
		Secrets: map[SecretId]Secret{},
		Crypter: NewCrypter(),
	}
}

func (r *memoryLocker) GetId() LockerId {
	return r.Id
}

// Add item to locker
func (r *memoryLocker) AddItem(
	secretName SecretId,
	key keys.Value,
	content PlainContent,
	resChan chan<- error) {
	//

	secret := r.Crypter.encrypt(keys.Value{}, content)
	r.Secrets[secretName] = secret

	//
	// Persist change
	//	go r.persistance.AddItem(clientId, lockerId, secretName, content)
	// return
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
	delete(r.Secrets, secretName)
	//
	// Persist change
	//	go r.persistance.RemoveItem(clientId, lockerId, secretName)
	// return
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
		// Check persistance
		close(resChan)
		return
	}
	resChan <- r.Crypter.decrypt(key, r.Secrets[secretName])
}
