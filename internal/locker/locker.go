package locker

import (
	"errors"

	"github.com/mes1234/golock/internal/client"
	"github.com/mes1234/golock/internal/keys"
)

// Locker is container for all secrect
type Locker struct {
	Crypter Crypter             // provide cryptgraphic functionality
	Id      LockerId            // Identifier of locker
	Client  client.ClientId     // Identifiers of all clients with access
	Secrets map[SecretId]Secret //Content of Locker

}

// Add item to locker
func (r *Locker) AddItem(
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
func (r *Locker) RemoveItem(
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
func (r *Locker) GetItem(
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
