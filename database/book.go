package database

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

func AddBook() error {
	client, err := mongodbClient()
	if err != nil {
		return err
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database(database).Collection(collectionBook)

	name, isbn := "fire & blood", "12345"

	st := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "Rasel",
		Age:  26,
	}

	result, err := coll.InsertOne(
		context.TODO(),
		bson.D{
			{"name", name},
			{"isbn", isbn},
			{"authors", st},
		})
	if err != nil {
		return err
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", string(jsonData))

	return nil
}
