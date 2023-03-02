package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Info struct {
	ID      int    `bson:"id"`
	Name    string `bson:"name"`
	Age     int    `bson:"age"`
	Address string `bson:"address"`
	Email   string `bson:"email"`
}

func main() {
	CLiOp := options.Client().ApplyURI("mongodb+srv://gopal476:gopal476@cluster0.jmfv2vw.mongodb.net/?retryWrites=true&w=majority")

	client, err := mongo.Connect(context.Background(), CLiOp)

	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	defer client.Disconnect(context.Background())

	database := client.Database("employee")
	collection := database.Collection("info")
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter Name :")
	scanner.Scan()
	name1 := scanner.Text()

	//checking in the database whether the name is there or not

	check := bson.M{"name": name1}
	fmt.Println(name1)
	var result *Info
	err = collection.FindOne(context.Background(), check).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	if result.ID != 0 {
		fmt.Println("User Found: ", name1)

		var newAge int
		fmt.Println("Enter new age: ")
		fmt.Scanln(&newAge)

		//can update  any other by taking input here

		update := bson.M{"$set": bson.M{"age": newAge}}
		_, err := collection.UpdateOne(context.Background(), check, update)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("User %s updated with age %d.\n", name1, newAge)
	} else {
		fmt.Printf("User not found with name %s.\n", name1)
	}

}
