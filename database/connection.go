package database

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

const (
	// export the connection string as MONGODB_URI
	mongodbURI = "MONGODB_URI"
	database   = "test"
	collection = "test-collection"
)

func Query() error {
	if err := godotenv.Load(); err != nil {
		fmt.Println("no .env file found")
	}
	uri := os.Getenv(mongodbURI)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database(database).Collection(collection)
	var result bson.M
	title := "book"

	err = coll.FindOne(context.TODO(), bson.D{{"name", title}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the name %s\n", title)
		return nil
	}
	if err != nil {
		return err
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", jsonData)

	return nil
}
