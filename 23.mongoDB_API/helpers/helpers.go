//? bson.D vs bson.M :: https://stackoverflow.com/questions/64281675/bson-d-vs-bson-m-for-find-queries
//? primitive.M vs bson.M

package helpers

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/DNahar74/golang-playlist/23.mongoDB_API/models"
)

var collection *mongo.Collection

// insertOneMovie is a function to insert a single document into a MongoDB collection
func insertOneMovie(movie models.Netflix) {
	result, err := collection.InsertOne(context.TODO(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted a single document: %+v\n", result)
}

// updateOneMovie is a function to update a single document in a MongoDB collection by ID
func updateOneMovie(movieID string)  {
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
	if err!= nil {
    log.Fatal(err)
  }
	defer cursor.Close(context.Background())

	var movies []bson.D

	for cursor.Next(context.Background()) {
		var movie bson.D
		err := cursor.Decode(&movie)
		if err!= nil {
      log.Fatal(err)
    }

		movies = append(movies, movie)
	}

	return movies
}