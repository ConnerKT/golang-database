package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type User struct {
	FullName string `json:"fullName"`
	Age      int    `json:"age"`
}

func readJSONFile(filename string) ([]User, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func main() {
	// Setting the connection string for MONGO
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connecting
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	//disconnect if error
	defer client.Disconnect(context.TODO())

	// Check the connection
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}
	fmt.Println("Connected to MongoDB!")

	// Access the database and collection if not create
	usersCollection := client.Database("testing").Collection("users")

	// Read users from JSON file
	users, err := readJSONFile("data.json")
	if err != nil {
		log.Fatal("Failed to read JSON file:", err)
	}

	// Prepare documents for insertion
	var documents []interface{}
	for _, user := range users {
		documents = append(documents, bson.D{
			{"fullName", user.FullName},
			{"age", user.Age},
		})
	}

	// Insert documents into MongoDB
	insertResult, err := usersCollection.InsertMany(context.TODO(), documents)
	if err != nil {
		log.Fatal("Failed to insert documents into MongoDB:", err)
	}

	// Print inserted IDs
	fmt.Println("Inserted IDs:", insertResult.InsertedIDs)
}
