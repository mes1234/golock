package persistance

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Retrieve collection to operate
func getDbCollection(collectionName string,
	dbName string) (
	collection *mongo.Collection,
	ctx context.Context) {
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

	collection = client.Database(dbName).Collection(collectionName)

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
