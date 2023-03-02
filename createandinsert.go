package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

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
	//Setting up mongoDB Clients

	CLiOp := options.Client().ApplyURI("mongodb+srv://gopal476:gopal476@cluster0.jmfv2vw.mongodb.net/?retryWrites=true&w=majority")

	// Here is mongodb connection
	client, err := mongo.Connect(context.Background(), CLiOp)
	if err != nil {
		fmt.Println("Error connecting to MongoDB.", err)
	}

	//disconnects from mongodb when the programs ends
	defer client.Disconnect(context.Background())

	/*
		//creating a database name  and collection

		startdb := client.Database("employee")
		pdcollection := startdb.Collection("info")
	*/

	collection := client.Database("employee").Collection("info")

	//read user input
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter Person Name: ")
	scanner.Scan()
	name := scanner.Text()

	var id int
	fmt.Println("Enter ID: ")
	fmt.Scanln(&id)

	var age int
	fmt.Println("Enter Person age: ")
	fmt.Scanln(&age)

	fmt.Println("Enter Person Address: ")
	scanner.Scan()
	address := scanner.Text()

	fmt.Println("Enter Person Email: ")
	scanner.Scan()
	email := scanner.Text()

	//creating a object

	Employee := Info{
		Id:      id,
		Name:    name,
		Age:     age,
		Address: address,
		Email:   email,
	}
	_, err = collection.InsertOne(context.Background(), Employee)
	if err != nil {
		fmt.Println("Error inserting documents:", err)
	}

	fmt.Println("The values you have entered is inserted sucessfully in MongoDB.")
}
