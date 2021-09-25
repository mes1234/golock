package persistance

import (
	"github.com/google/uuid"
	"github.com/mes1234/golock/adapters"
)

type SecretRepository interface {
	Insert(*SecretPersisted) error                          //Insert secret data to DB and assign ID
	Retrieve(lockerId uuid.UUID) ([]SecretPersisted, error) // Get secret from DB
}

type secretRepository struct{}

func NewSecretRepository() SecretRepository {
	return &secretRepository{}
}

func (sr secretRepository) Insert(secretDetails *SecretPersisted) (err error) {

	collection, ctx := getDbCollection("secrets", "secrets")

	document, err := convertToMongoDocument(secretDetails)
	if err != nil {
		return
	}

	collection.InsertOne(ctx, document)
	err = nil

	collection.Database().Client().Disconnect(ctx)
	return
}

func (sr secretRepository) Retrieve(lockerId uuid.UUID) (secrets []SecretPersisted, err error) {

	collection, ctx := getDbCollection("secrets", "secrets")

	document, err := convertToMongoDocument(adapters.LockerId{
		LockerId: lockerId,
	})
	if err != nil {
		return
	}

	cursor, err := collection.Find(ctx, document)

	if err != nil {
		return
	}
	secrets = make([]SecretPersisted, 0)
	if err = cursor.All(ctx, &secrets); err != nil {
		return
	}
	err = nil
	return
}
