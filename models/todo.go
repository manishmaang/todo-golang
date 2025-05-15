package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	// primitive is the package being imported , and it provides the mongodb data types for our schema to store data
)

type Todo struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Title     string             `bson:"title" json:"title"`
	Description string           `bson:"description" json:"description"`
	Completed bool               `bson:"completed" json:"completed"`
	User_Id   primitive.ObjectID  `bson:"user_id"  json:"user_id"`
}
