package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	// primitive is the package being imported , and it provides the mongodb data types for our schema to store data
)

type Users struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Username     string             `bson:"username" json:"username"`
	Password     string           `bson:"password" json:"password"`
	Tasks        []primitive.ObjectID  `bson:"tasks" json:"tasks"`
}
