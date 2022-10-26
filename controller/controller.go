package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Yashsharma1911/mongoapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "{get connection link from mongodb atlas}"
const dbName = "netflix"
const collectionName = "watchlist"

// reference to the database collection
var collection *mongo.Collection

func init() {

	clientOption := options.Client().ApplyURI(connectionString)

	// connecting to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDb connection successful 🎉")

	// references to collection
	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("Collection instance is ready 🚀")
}

// Insert one movie
func insertOneMovie(movie model.Netflix) {
	insertResult, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

// Update movie watched status
func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count ", result.ModifiedCount)
}

// Delete one movie
func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("deleteCount ", deleteCount)
}

// Delete all movies
func deleteAllMovie() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of movies deleted ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// Get all movies in the database
func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	defer cur.Close(context.Background())

	var movies []primitive.M

	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	return movies
}

// Actual controllers
func HomeServe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("<h1>Hey this is home</h1>")
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)

	// sending message on successful creation
	var message model.SendMessage
	message.Message = "Successfully Added"
	message.DataAdded = &movie
	json.NewEncoder(w).Encode(message)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	// get id from url
	params := mux.Vars(r)
	id := params["id"]
	updateOneMovie(id)

	// sending data on successful
	var message model.SendMessage
	message.Message = "Successfully updated"
	json.NewEncoder(w).Encode(message)
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	// get id from url
	params := mux.Vars(r)
	id := params["id"]
	deleteOneMovie(id)

	// sending data on successful
	var message model.SendMessage
	message.Message = "Successfully deleted"
	json.NewEncoder(w).Encode(message)
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}
