package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options" // The correct import
)

const connection_string = "mongodb://localhost:27017/"
const db_name = "golang-todo"

var DB *mongo.Database
func init() { //special functions jo ki sirf pure application me ek baar run hote hai, that to jb application start ho rha hota hai

	//client option
	clientOption := options.Client().ApplyURI(connection_string).SetMaxPoolSize(10) 
	// options is the package imported in line 9, options.client() is a function which creates a setting struct(obj) which has connection configuration

	//connection with mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)
	// (context.todo)so it basically defines some setting for each db operation but context.todo is like empty setting where nothing is handled but will be later on

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Conneceted with the mongoDB");

    DB = client.Database(db_name)
}

func GetDB() *mongo.Database{
	return DB
}
