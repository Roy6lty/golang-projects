package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	model "github.com/olowoleru/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = `mongodb://admin:password@0.0.0.0:27017`
const dbName = "netflix"
const collectionName = "watchlist"

//Most Important

var collection *mongo.Collection

// connect with mongodb

func init(){
	//initialization of the Mongodb
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Mongodb connection established")

	collection = client.Database(dbName).Collection(collectionName)
}

// MongoDB helpers -file

//insert 1 record
func insertOneMovie(movie model.Netflix){
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id:", inserted.InsertedID)

}

// update 1 record
func updateOneMovie(movieId string){
	//create an ID from Hex.....(Create an Bson ID for mongoDB)
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set":bson.M{"watched":true}}


	//Updating a value
	//1. first you filter for the value you want
	//2. Then update tat value 
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	// Number of documents updated / modified
	fmt.Println("modified count: ", result.ModifiedCount)
}

func deleteOneMovie(movieID string){
	id, _ := primitive.ObjectIDFromHex(movieID) //convert sring to bson
	filter := bson.M{"_id": id}
	deleteCount , err :=collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie got deleted successfully: ", deleteCount) 
}

//
func deleteAllMovies() int64 {
	// empty filter means select everything as everything matches
	deleteResult, err  := collection.DeleteMany(context.Background(),  bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of movies delete: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

//get all movies from database

func getAllMovies()[]primitive.M{
cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	//cur.Next(context.Background()) returns a bool
	// indicating if a next value exist
	// so the statement can be interpreted as 
	// for True... do the loop

	for cur.Next(context.Background()) {
		var movie bson.M
		// because passing the movie by reference no need to
		// assign the movie variable after
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)

	}
	defer cur.Close(context.Background())
	return movies

}

//controllers
func GetMyAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlcode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlcode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlcode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params)
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlcode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlcode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count :=deleteAllMovies()
	json.NewEncoder(w).Encode(count)

}

