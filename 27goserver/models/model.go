package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type MongoUser struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Phone_number string             `json:phone_number" bson:"phone_number"`
	Name         string             "json:name"
	Message      string             "json:message"
}
