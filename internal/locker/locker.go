package locker

import (
	"errors"

	"github.com/mes1234/golock/internal/client"
	"github.com/mes1234/golock/internal/keys"
)

// Locker is container for all secrect
type Locker struct {
	crypter Crypter             // provide cryptgraphic functionality
	Id      LockerId            // Identifier of locker
	Client  client.ClientId     // Identifiers of all clients with access
	Secrets map[SecretId]Secret //Content of Locker

}

func (r *Locker) AddItem(
	secretName SecretId,
	key keys.Value,
	content PlainContent,
	resChan chan<- error) {
	//

	secret := r.crypter.encrypt(keys.Value{}, content)
	r.Secrets[secretName] = secret

	//
	// Persist change
	//	go r.persistance.AddItem(clientId, lockerId, secretName, content)
	// return
	resChan <- nil
}

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

func (r *Locker) GetItem(
	lockerId LockerId,
	key keys.Value,
	secretName SecretId,
	resChan chan<- struct {
		PlainContent
		error
	}) {
	// Ensure thread safety
	//
	if _, ok := r.Secrets[secretName]; !ok {
		// Check persistance
		resChan <- struct {
			PlainContent
			error
		}{PlainContent{}, errors.New("no item found for given secret id")}
	}
	resChan <- struct {
		PlainContent
		error
	}{r.crypter.decrypt(key, r.Secrets[secretName]), nil}
}
