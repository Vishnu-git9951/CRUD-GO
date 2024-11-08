package main

import (
	db "awesomeProject1/dbConnectPack"
	"fmt"
)

func main() {
	// Initialize MongoDB connection using mgm
	db.Connect()

	// Retrieve and print the first document in the 'user' collection
	firstUser, err := db.GetFirstDocument()
	if err != nil {
		fmt.Printf("Error retrieving first documents: %v\n", err)
		return
	}

	if firstUser != nil {
		fmt.Printf("First user document: %+v\n\n", firstUser.Name)
	} else {
		fmt.Println("No documents found in the 'user' collection.")
	}

	// Get and print the total document count in the 'user' collection
	totalCount, err := db.GetTotalDocumentCount("user")
	if err != nil {
		fmt.Printf("Error counting documents: %v\n", err)
		return
	}
	fmt.Printf("Total document count in 'user' collection: %d\n", totalCount)
}
