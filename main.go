package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	//Setting my atlas db to a variable to reuse later (env to be made)
	var atlasDb = "mongodb+srv://golang_user:password1234@cluster0.nnbdzbq.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
	fmt.Println(atlasDb)
	//Anytime you make requests to a server (our db),
	//you should create a context using context.TODO()
	//that the server will accept
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	usersCollection := client.Database("testing").Collection("users")
	// insert a single document into a collection
	// create a bson.D object
	// user := bson.D{{"fullName", "User 1"}, {"age", 30}}
	// insert the bson object using InsertOne()
	// result, err := usersCollection.InsertOne(context.TODO(), user)
	// check for errors in the insertion
	// if err != nil {
	// 	panic(err)
	// }
	// // display the id of the newly inserted object
	// fmt.Println(result.InsertedID)

	// insert multiple documents into a collection
	// create a slice of bson.D objects
	users := []interface{}{
		bson.D{{"fullName", "User 2"}, {"age", 25}},
		bson.D{{"fullName", "User 3"}, {"age", 20}},
		bson.D{{"fullName", "User 4"}, {"age", 28}},
	}
	// insert the bson object slice using InsertMany()
	results, err := usersCollection.InsertMany(context.TODO(), users)
	// check for errors in the insertion
	if err != nil {
		panic(err)
	}
	// display the ids of the newly inserted objects
	fmt.Println(results.InsertedIDs)
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {

	}
	if err != nil {
		panic(err)
	}
}
