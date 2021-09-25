package locker

import (
	"github.com/mes1234/golock/internal/keys"
	"log"

	"github.com/google/uuid"
	"github.com/mes1234/golock/internal/client"
	"github.com/mes1234/golock/persistance"
)

type dbRepository struct {
	ClientId uuid.UUID
}

func getDbRepository(clientId client.Id) Repository {
	return &dbRepository{
		ClientId: clientId,
	}
}

func (r *dbRepository) Update(l Locker, lockerId uuid.UUID, resChan chan<- bool) {
	log.Print("update to db")

	dbAccess := persistance.NewSecretRepository()

	items := l.ItemsToCommit()

	for k, v := range items {
		itemToPersist := persistance.SecretPersisted{
			Active:     v.Active,
			Revision:   v.Revision,
			Content:    v.Content,
			SecretName: k,
			LockerId:   lockerId,
			ClientId:   l.GetClientId(),
		}
		err := dbAccess.Insert(&itemToPersist)
		if err != nil {
			resChan <- false
		}
	}

	resChan <- true
}

func (r *dbRepository) Get(lockerId uuid.UUID, resChan chan<- Locker) {

	dbAccess := persistance.NewSecretRepository()

	secrets, err := dbAccess.Retrieve(lockerId)
	if err != nil {
		close(resChan)
	}
	newLocker := GetMemoryLocker(r.ClientId, lockerId)

	for _, s := range secrets {
		err := make(chan error, 0)
		go newLocker.AddItem(s.SecretName, keys.Value{}, s.Content, s.Revision, err)
		_ = <-err
	}
	resChan <- newLocker
}

// Create in Db implementation single item is not represented by locker
func (r *dbRepository) Create(lockerId uuid.UUID, resChan chan<- uuid.UUID) {
	resChan <- lockerId
}
