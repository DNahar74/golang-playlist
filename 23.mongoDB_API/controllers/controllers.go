//todo: (1) Understand context and it's necessity

package controllers

import (
	"context"
	"fmt"
	"log"

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
