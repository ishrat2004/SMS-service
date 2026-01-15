package main

import (
	"context"
	"encoding/json"
	"fmt"
	"kafka/controller"
	"kafka/model"

	"github.com/segmentio/kafka-go"
)

type User struct {
	Phone_number string "json:phone_number"
	Name         string "json:name"
	Message      string "json:message"
}

func consume(ctx context.Context) {
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	broker1Address := "localhost:9092"

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker1Address},
		Topic:   "letmeeshoyou2k26",
		GroupID: "go-consumer-group",
	})

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic(err)
		}

		fmt.Println("Raw Kafka Message:", string(msg.Value))

		user := User{}

		err = json.Unmarshal(msg.Value, &user)
		if err != nil {
			panic(err)
		}
		mongoUser := model.MongoUser{
			Phone_number: user.Phone_number,
			Name:         user.Name,
			Message:      user.Message,
		}
		controller.AddMessage(mongoUser)
		fmt.Printf("User object: %+v\n", user)
		fmt.Println(user.Phone_number)
		fmt.Println("Name" + user.Name)
		fmt.Println("Message" + user.Message)

	}
}

func main() {
	fmt.Println("testing kafka consumer")
	ctx := context.Background()
	consume(ctx)

}
