package persistance

import (
	"context"
	"log"
	"time"

	"github.com/mes1234/golock/adapters"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type ClientRepository interface {
	Insert(*adapters.Client) error   //Insert client data to DB and assing ID
	Retrieve(*adapters.Client) error // Get id for given username and password if matched
}

type clientRepository struct{}

func NewClientRepository() ClientRepository {
	return &clientRepository{}
}

func (cr clientRepository) Insert(clientDetails *adapters.Client) (err error) {

	collection, ctx := getDbCollection()

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

	collection, ctx := getDbCollection()

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

// Retrieve collection to operate
func getDbCollection() (collection *mongo.Collection, ctx context.Context) {
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://root:example@localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("clients").Collection("clients")

	return

}

// convert any object to mongo document
func convertToMongoDocument(v interface{}) (doc *bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}
	err = bson.Unmarshal(data, &doc)
	return
}
