package persistance

import (
	"log"

	"github.com/mes1234/golock/adapters"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

type ClientRepository interface {
	Insert(*adapters.Client) error   //Insert client data to DB and assign ID
	Retrieve(*adapters.Client) error // Get id for given username and password if matched
}

type clientRepository struct{}

func NewClientRepository() ClientRepository {
	return &clientRepository{}
}

func (cr clientRepository) Insert(clientDetails *adapters.Client) (err error) {

	collection, ctx := getDbCollection("clients", "clients")

	document, err := convertToMongoDocument(clientDetails)
	if err != nil {
		return
	}

	collection.InsertOne(ctx, document)
	err = nil

	collection.Database().Client().Disconnect(ctx)
	return
}

func (cr clientRepository) Retrieve(clientDetails *adapters.Client) (err error) {

	collection, ctx := getDbCollection("clients", "clients")

	var data interface{}

	if clientDetails.ClientName != "" {
		data, _ = convertToMongoDocument(adapters.ClientName{
			ClientName: clientDetails.ClientName,
		})
	} else {
		data, _ = convertToMongoDocument(adapters.ClientId{
			ClientId: clientDetails.ClientId,
		})
	}

	filterCursor, err := collection.Find(ctx, data)
	if err != nil {
		log.Fatal(err)
	}

	filterCursor.Next(ctx)

	current := filterCursor.Current

	bson.Unmarshal(current, &clientDetails)

	collection.Database().Client().Disconnect(ctx)

	err = nil
	return
}
