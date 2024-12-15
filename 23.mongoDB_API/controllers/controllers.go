//todo: (1) Understand context and it's necessity

package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/DNahar74/golang-playlist/23.mongoDB_API/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://nahardarsh55:<db-password>@go-playlist.vhenj.mongodb.net/?retryWrites=true&w=majority&appName=go-playlist"
const dbName = "NETFLIX"
const collectionName = "watchlist"

var collection *mongo.Collection

// Connect with MongoDB

// init only runs for initialization (only once)
func init() {
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connected successfully")

	collection = client.Database(dbName).Collection(collectionName)

	// If the collection reference exists
	fmt.Println("Collection reference is available")
}

// insertOneMovie is a function to insert a single document into a MongoDB collection
func insertOneMovie(movie models.Netflix) {
	result, err := collection.InsertOne(context.TODO(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted a single document: %+v\n", result)
}

// updateOneMovie is a function to update a single document in a MongoDB collection by ID
func updateOneMovie(movieID string) {
	ObjectID, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": ObjectID}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("update one result: %+v\n", result)
}

// deleteOneMovie is a function to delete a single document in a MongoDB collection by ID
func deleteOneMovie(movieID string) {
	ObjectID, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": ObjectID}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("delete one result: %+v\n", result)
}

// deleteAllMovies is a function to delete all documents in a MongoDB collection
func deleteAllMovies() {
	result, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("delete all result: %+v\n", result)
}

// getAllMovies is a function to get all documents in a MongoDB collection and return them as a slice of bson.D
func getAllMovies() []bson.D {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	var movies []bson.D

	for cursor.Next(context.Background()) {
		var movie bson.D
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	return movies
}

// This is controllers.go part

// GetAllMovies writes the data of all movies to the writer
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	allMovies := getAllMovies()

	// Craft the response to be user-friendly
	json.NewEncoder(w).Encode(allMovies)
}

// CreateMovie inserts a new movie in the database
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "PATCH")

	var movie models.Netflix

	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		log.Fatal(err)
	}

	insertOneMovie(movie)
	w.Write([]byte("Data inserted successfully"))
}

// MarkAsWatched marked the movie as watched in the database
func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])

	w.Write([]byte("Marked the movie as watched"))
}

// DeleteOneMovie deletes a movie from the database by its ID
func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	deleteOneMovie(params["id"])
	w.Write([]byte("This movie has been deleted"))
}

// DeleteAllMovies deletes all movies from the database
func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	deleteAllMovies()
	w.Write([]byte("All movies have been deleted"))
}
