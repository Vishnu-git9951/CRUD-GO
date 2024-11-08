package dbConnectPack

import (
	"fmt"
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect to MongoDB and initialize the Employee database connection using mgm
func Connect() {
	err := mgm.SetDefaultConfig(nil, "Employee", options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	fmt.Println("Connected to MongoDB with mgm and selected database: Employee")
}

// User represents a document in the 'user' collection
type User struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Email            string `json:"email" bson:"email"`
}

// GetFirstDocument retrieves the first document from the 'user' collection
func GetFirstDocument() (*User, error) {
	user := &User{}
	err := mgm.Coll(user).First(bson.M{}, user)
	println(err, "errore")
	if err != nil {
		if err.Error() == mongo.ErrNoDocuments.Error() {
			fmt.Println("No documents found in the 'user' collection")
			return nil, nil
		}
		return nil, fmt.Errorf("failed to retrieve first document: %v", err)
	}
	return user, nil
}

// GetTotalDocumentCount returns the total number of documents in the specified collection
func GetTotalDocumentCount(collectionName string) (int64, error) {
	coll := mgm.Coll(&User{})
	filter := bson.M{
		"name": "abhinav singh", // Match documents where the "name" field is exactly "Alice"
	}
	count, err := coll.CountDocuments(mgm.Ctx(), filter)
	if err != nil {
		return 0, fmt.Errorf("failed to count documents in collection %s: %v", collectionName, err)
	}
	return count, nil
}
