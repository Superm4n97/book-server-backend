package database

import (
	"context"
	"encoding/json"
	"github.com/Superm4n97/book-server-backend/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func resultToBook(result bson.M) (*utils.Book, error) {
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		return nil, err
	}
	var book utils.Book
	err = json.Unmarshal(jsonData, &book)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func AddBook(book *utils.Book) error {
	return query(collectionBook, func(coll *mongo.Collection) error {
		_, err := coll.InsertOne(
			context.TODO(),
			bson.D{
				{"name", book.Name},
				{"genre", book.Genre},
				{"authors", book.Authors},
			})
		return err
	})
}

func GetBook(name string) (*utils.Book, error) {
	var result bson.M
	err := query(collectionBook, func(coll *mongo.Collection) error {
		err := coll.FindOne(
			context.TODO(),
			bson.D{
				{"name", name},
			}).Decode(&result)
		return err
	})
	if err != nil {
		return nil, err
	}
	return resultToBook(result)
}

func DeleteBook(name string) error {
	return query(collectionBook, func(coll *mongo.Collection) error {
		_, err := coll.DeleteOne(
			context.TODO(),
			bson.D{
				{"name", name},
			})
		return err
	})
}
