package main

import (
	"fmt"
	"log"
	"time"
)

type Example struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

func main() {

	cluster := couchInit()

	fmt.Println("\nConnected to the bucket")

	example := Example{
		Name: "Example 1",
		Date: time.Now().Format(time.RFC3339),
	}

	key := "example-1"

	// Insert data

	err := cluster.insertData(key, example)

	if err != nil {

		log.Fatalf("Error inserting data: %s\n", err.Error())

	}

	fmt.Println("\nCreate document success")

	// Get data

	doc, err := cluster.getData(key)

	if err != nil {

		log.Fatalf("Error getting data:%s\n", err.Error())

	}

	fmt.Printf("Data is: %v\n", doc)

	// Update Data

	doc.Name = "Example 2"

	err = cluster.updateData(key, doc)

	if err != nil {

		log.Fatalf("Error updating the document: %s\n", err.Error())
	}

	fmt.Println("Updated successfully")

	//Delete data

	err = cluster.deleteData(key)

	if err != nil {

		log.Fatalf("Error deleting data: %s\n", err.Error())

	}

	fmt.Println("Deleted successfully")

}
