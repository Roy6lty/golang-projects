package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = `mongodb://admin:pasword@0.0.0.0:27017`
const databaseName = "Ecommerce"

// const dbName = "netflix"
// const collectionName = "watchlist"

//Most Important

// connect with mongodb
func DBSetup() *mongo.Client {
	//initialization of the Mongodb
	clientOption := options.Client().ApplyURI(connectionString)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	//connect to mongodb
	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		log.Fatal(err)
	}

	defer cancel()

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("Failed to connect to mongo server :()")
	}
	fmt.Println("Mongodb connection established")

	return client

}

var Client *mongo.Client = DBSetup()

func UserData(client *mongo.Client, collectionName string) *mongo.Collection {
	//Create a new collection for users
	var userCollection *mongo.Collection = client.Database(databaseName).Collection(collectionName)
	return userCollection

}

func ProductData(client *mongo.Client, collectionName string) *mongo.Collection {
	var productCollection *mongo.Collection = client.Database(databaseName).Collection(collectionName)
	return productCollection

}
