package database

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Superm4n97/book-server-backend/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func resultToAuthor(result bson.M) (*utils.Author, error) {
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		return nil, err
	}
	fmt.Println("jsonData: ", string(jsonData))

	var athr utils.Author
	err = json.Unmarshal(jsonData, &athr)
	if err != nil {
		return nil, err
	}

	return &athr, nil
}

func AddAuthor(author utils.Author) error {
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

func GetAuthor(name string) (*utils.Author, error) {
	client, err := mongodbClient()
	if err != nil {
		return nil, err
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
		return nil, err
	}
	return resultToAuthor(result)
}
