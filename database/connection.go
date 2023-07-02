package database

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

const (
	// export the connection string as MONGODB_URI
	mongodbURI       = "MONGODB_URI"
	database         = "book-server"
	collectionBook   = "book"
	collectionAuthor = "author"
)

func mongodbClient() (*mongo.Client, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("no .env file found")
	}
	uri := os.Getenv(mongodbURI)

	c, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return c, nil
}

type helper func(coll *mongo.Collection) error

func query(helper helper) error {
	client, err := mongodbClient()
	if err != nil {
		return err
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database(database).Collection(collectionAuthor)

	return helper(coll)
}
