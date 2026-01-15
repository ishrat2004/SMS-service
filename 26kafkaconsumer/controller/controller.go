package controller

import (
	"context"
	"fmt"
	"kafka/model"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connection_string = "mongodb+srv://ishratalikhan004:Ishratkhan22@cluster0.fzscbb3.mongodb.net/?appName=Cluster0"

const dbName = "sms_messages"

const colName = "all_messages"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connection_string)
	fmt.Println("client connection")

	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		panic(err)
	}
	fmt.Println("connected to mongodb")
	collection = client.Database(dbName).Collection(colName)

	fmt.Println("collection ref is ready")
}

func insertMessage(user model.MongoUser) {
	fmt.Println("into the function of adding ")
	inserted, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	fmt.Println("inserted is ", inserted)
	fmt.Println("inserted one movie", inserted.InsertedID)
}

func AddMessage(user model.MongoUser) {
	fmt.Println("into adding movies route ")
	//    w.Header().Set("Content-type", "applications/json")
	insertMessage(user)
}
