package mongo

import (
	"context"
	"github.com/akraskovski/go-delivery-service/internal/app/model"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

const collectionName string = "orders"

type OrderRepository struct {
	store *Store
}

func (repository *OrderRepository) Create(order *model.Order) (*model.Order, error) {
	if _, err := repository.store.Collection(collectionName).InsertOne(context.Background(), order); err != nil {
		log.Fatal(err)
	}
	return order, nil
}

func (repository *OrderRepository) FindAll() ([]*model.Order, error) {
	var results []*model.Order
	cursor, err := repository.store.Collection(collectionName).Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {

		// create a value into which the single document can be decoded
		var elem model.Order
		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	return results, nil
}
