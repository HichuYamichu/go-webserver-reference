package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User : represents user
type User struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name string             `bson:"name,omitempty" json:"name,omitempty"`
	Age  int                `bson:"age,omitempty" json:"age,omitempty"`
}
