package database

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Superm4n97/book-server-backend/utils/author"
	"go.mongodb.org/mongo-driver/bson"
)

func AddAuthor(author author.Author) error {
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

	_, err = coll.InsertOne(
		context.TODO(),
		bson.D{
			{"name", author.Name},
			{"email", author.Email},
		})
	if err != nil {
		return err
	}

	return nil
}

func GetAuthor(name string) error {
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

	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{"name", name}}).Decode(&result)
	if err != nil {
		return err
	}

	fmt.Println("test: ")

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
	return nil
}
