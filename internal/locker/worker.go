package locker

import (
	"github.com/google/uuid"
	"github.com/mes1234/golock/internal/client"
)

// Attach to locker to obtain req and response channels
// TODO change it to allow generic boxed version of request & response
func Attach() (requestCh chan<- client.ClientId,
	responseCh <-chan LockerId) {

	clientCh := make(chan client.ClientId)
	lockerCh := make(chan LockerId)

	requestCh = clientCh
	responseCh = lockerCh

	go handleNewLocker(clientCh, lockerCh)

	return

}

func handleNewLocker(
	clientCh <-chan client.ClientId,
	lockerCh chan<- LockerId) {

	clientId := <-clientCh

	newLocker := Locker{
		Client:  clientId,
		Id:      uuid.New(),
		Secrets: make(map[SecretId]Secret),
	}
	memoryRepository[newLocker.Id] = newLocker
	lockerCh <- newLocker.Id

}
