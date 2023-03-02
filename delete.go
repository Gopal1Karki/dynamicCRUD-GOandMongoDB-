package main

import (
	"context"
	"fmt"
	"log"
	"strings"

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
		fmt.Println("Error connecting MongoDB database!!", err)
	}

	defer client.Disconnect(context.Background())

	collection := client.Database("employee").Collection("info")

	//scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter Id which you want to delete: ")
	//scanner.Scan()
	//Id := scanner.Text()
	var Id int
	fmt.Scanln(&Id)
	//Checking the id in the database

	check := bson.M{"id": Id}
	var result *Info

	err = collection.FindOne(context.Background(), check).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	if result.ID != 0 {
		fmt.Println("User found with id: ", Id)
		fmt.Println(result)

		//confirming again wheher to delete data or not
		fmt.Printf("\n")
		var inp string
		fmt.Println("Are you sure you want to delete ??? (Y/N): ")
		fmt.Scanln(&inp)
		inp1 := strings.ToLower(inp)

		if inp1 == "y" {
			_, err := collection.DeleteOne(context.Background(), check)
			if err != nil {
				fmt.Println("Error deleting data from database! ", err)
			}
			fmt.Printf("Record is deleted of Id: %d ", Id)
		} else {
			fmt.Println("Employee deletion Cancelled by the User")
		}
	} else {
		fmt.Printf("Employee not found with ID: ", Id)
	}

}
