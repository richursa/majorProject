package db

import (
	"context"
	"log"

	"../blockchain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database = "blockchain"
var mongoCollection = "athena"

func GetClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func ReturnBlockFromDB(client *mongo.Client, filter bson.M) blockchain.Block {
	var b blockchain.Block
	collection := client.Database(database).Collection(mongoCollection)
	documentReturned := collection.FindOne(context.TODO(), filter)
	documentReturned.Decode(&b)
	return b
}
