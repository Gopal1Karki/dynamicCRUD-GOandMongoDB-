package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Info struct {
	Id      int
	Name    string
	Age     int
	Address string
	Email   string
}

func main() {
	// Connect to MongoDB
	CLiOp := options.Client().ApplyURI("mongodb+srv://gopal476:gopal476@cluster0.jmfv2vw.mongodb.net/?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.Background(), CLiOp)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer client.Disconnect(context.Background())

	// Select the database and collection to use
	collection := client.Database("employee").Collection("info")

	// Find all documents in the collection
	cursor, err := collection.Find(context.Background(), bson.D{})

	if err != nil {
		fmt.Println("Error finding documents:", err)
		return
	}
	defer cursor.Close(context.Background())

	// Iterate over the documents and print their fields
	for cursor.Next(context.Background()) {
		var info Info
		err := cursor.Decode(&info)
		if err != nil {
			fmt.Println("Error decoding document:", err)
		}
		fmt.Println("Id", info.Id)
		fmt.Println("Name:", info.Name)
		fmt.Println("Age:", info.Age)
		fmt.Println("Address:", info.Address)
		fmt.Println("Email:", info.Email)
		fmt.Println()
	}

	// Check for errors during iteration
	if err := cursor.Err(); err != nil {
		fmt.Println("Error iterating over documents:", err)
		return
	}
}
