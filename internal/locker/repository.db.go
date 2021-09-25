package locker

import (
	"log"

	"github.com/google/uuid"
	"github.com/mes1234/golock/internal/client"
	"github.com/mes1234/golock/persistance"
)

type dbRepository struct {
}

func getDbRepository(clientId client.Id) Repository {

	return &dbRepository{}
}

func (r *dbRepository) UpdateLocker(l Locker, lockerId uuid.UUID, resChan chan<- bool) {
	log.Print("update to db")

	dbAccess := persistance.NewSecretRepository()

	items := l.ItemsToCommit()

	for k, v := range items {
		itemToPersisit := persistance.SecretPersisted{
			Active:     v.Active,
			Revision:   v.Revision,
			Content:    v.Content,
			SecretName: k,
			LockerId:   lockerId,
			ClientId:   l.GetClientId(),
		}
		dbAccess.Insert(&itemToPersisit)
	}

	resChan <- true
}

func (r *dbRepository) GetLocker(lockerId uuid.UUID, resChan chan<- Locker) {
	close(resChan)
}

func (r *dbRepository) InitLocker(lockerId uuid.UUID, resChan chan<- uuid.UUID) {
	resChan <- lockerId
}
