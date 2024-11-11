package database

import (
	"context"
	"fmt"
	"log"
	model "olowoleru/pams/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = `mongodb://admin:password@0.0.0.0:27017`
const dbName = "pams"
const collectionName = "inventory"

var collection *mongo.Collection

func init() {
	//initialization of the Mongodb
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongodb connection established")

	collection = client.Database(dbName).Collection(collectionName)
}

type CatalogueType interface {
	model.CatalogueSchema | model.BuildingCatalogue | model.VehicleCatalogue | model.EquipmentCatalogue
}

// insert One
// helpers
func InsertOneCatalogue(catalogue interface{}) {
	inserted, err := collection.InsertOne(context.Background(), catalogue)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id:", inserted.InsertedID)

}

// deleteOne
func DeleteOneCatalogueByID[T CatalogueType](catalogueID string) {
	id, err := primitive.ObjectIDFromHex(catalogueID)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id} // filter for catalogue ID
	collection.DeleteOne(context.Background(), filter)
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie got deleted successfully: ", deleteCount)

}

func GetOneCatalogueByID(catalogueID string) {
	id, err := primitive.ObjectIDFromHex(catalogueID)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id} // filter for catalogue ID
	var catalogue bson.M

	result := collection.FindOne(context.Background(), filter)
	result.Decode(&catalogue) //pass the result into catalogue
	if err != nil {
		log.Fatal(err)
	}

	return
}
