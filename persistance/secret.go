package persistance

type SecretRepository interface {
	Insert(*SecretPersisted) error   //Insert secret data to DB and assing ID
	Retrieve(*SecretPersisted) error // Get secret from DB
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

func (sr secretRepository) Retrieve(secretDetails *SecretPersisted) (err error) {

	// TODO  implement this

	err = nil
	return
}
