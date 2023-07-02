package database

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Superm4n97/book-server-backend/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func resultToAuthor(result bson.M) (*utils.Author, error) {
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		return nil, err
	}
	var athr utils.Author
	err = json.Unmarshal(jsonData, &athr)
	if err != nil {
		return nil, err
	}

	return &athr, nil
}

func AddAuthor(author *utils.Author) error {
	return query(collectionAuthor, func(coll *mongo.Collection) error {
		_, err := coll.InsertOne(
			context.TODO(),
			bson.D{
				{"name", author.Name},
				{"email", author.Email},
			})
		return err
	})
}

func GetAuthor(name string) (*utils.Author, error) {
	var result bson.M
	var err = query(collectionAuthor, func(coll *mongo.Collection) error {
		err := coll.FindOne(context.TODO(), bson.D{{"name", name}}).Decode(&result)
		return err
	})
	if err != nil {
		return nil, err
	}
	return resultToAuthor(result)
}

func DeleteAuthor(name string) error {
	return query(collectionAuthor, func(coll *mongo.Collection) error {
		_, err := coll.DeleteOne(
			context.TODO(),
			bson.D{
				{"name", name},
			})
		return err
	})
}

func ListAuthor() (*utils.AuthorList, error) {
	var cursor *mongo.Cursor
	err := query(collectionAuthor, func(coll *mongo.Collection) error {
		var err error
		cursor, err = coll.Find(
			context.TODO(),
			bson.D{
				{"", ""}, //can not list all authors in this way
			})
		return err
	})
	if err != nil {
		return nil, err
	}

	var result bson.A
	_ = cursor.All(context.TODO(), &result)
	fmt.Println("------------------------cursor--------------------")
	fmt.Println(result)

	return nil, nil
}
