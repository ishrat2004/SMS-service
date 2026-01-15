package controller

import (
	"context"
	"encoding/json"
	"fmt"
	model "goserver/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connection_string = "mongodb+srv://ishratalikhan004:Ishratkhan22@cluster0.fzscbb3.mongodb.net/?appName=Cluster0"

const dbName = "sms_messages"
const colName = "all_messages"

var collection *mongo.Collection

func init() {

	clientOption := options.Client().ApplyURI(connection_string)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		fmt.Println("getting error")
		fmt.Println(err)
		panic(err)
	}
	collection = client.Database(dbName).Collection(colName)

	fmt.Println("collection ref is ready")

}

// func getMessage(phone_number int64) []model.MongoUser {
// 	fmt.Println("inside the function")
// 	curr, err := collection.Find(context.Background(), bson.M{"phone_number": phone_number})
// 	if err != nil {
// 		fmt.Println(err)
// 		log.Fatal(err)
// 	}

// 	var messages []model.MongoUser
// 	for curr.Next(context.Background()) {
// 		var message model.MongoUser
// 		err := curr.Decode(&message)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		messages = append(messages, message)
// 	}
// 	fmt.Println("messages are ", messages)
// 	defer curr.Close(context.Background())
// 	return messages
// }

func getMessage(phone_number string) []model.MongoUser {
	temp := bson.D{{"phone_number", phone_number}}
	curr, err := collection.Find(
		context.Background(),
		temp,
	)
	if err != nil {
		log.Println("Finding error:", err)
	}
	defer curr.Close(context.Background())
	var messages []model.MongoUser
	for curr.Next(context.Background()) {
		var message model.MongoUser
		err := curr.Decode(&message)
		if err != nil {
			log.Fatal(err)
		}
		messages = append(messages, message)
	}
	return messages

}

func getAllMessages() []primitive.M {
	fmt.Println("inside the function")
	curr, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer curr.Close(context.Background())
	var messages []primitive.M
	for curr.Next(context.Background()) {
		var message bson.M
		err := curr.Decode(&message)
		if err != nil {
			log.Fatal(err)
		}
		messages = append(messages, message)
	}
	return messages

}

func GetAllMessages(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside the get all messages route")
	w.Header().Set("Content-type", "applications/json")
	allMessages := getAllMessages()
	fmt.Println("all messages are ", allMessages)
	json.NewEncoder(w).Encode(allMessages)
	return

}
func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside the get user route")
	vars := mux.Vars(r)
	string_phone_number := vars["phoneNumber"]
	// fmt.Println("phone number typ")
	w.Header().Set("Content-type", "applications/json")
	user_models := getMessage(string_phone_number)
	// fmt.Println("user is ", user_models)
	var only_mesages []string
	for _, message := range user_models {
		only_mesages = append(only_mesages, message.Message)
	}
	json.NewEncoder(w).Encode(only_mesages)
	return
}
