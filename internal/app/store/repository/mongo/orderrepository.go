package mongo

import (
	"context"
	"fmt"
	"github.com/akraskovski/go-delivery-service/internal/app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const collectionName string = "orders"

type OrderRepository struct {
	store *Store
}

func (repository *OrderRepository) Create(order *model.Order) (string, error) {
	order.Id = primitive.NewObjectID()
	res, err := repository.store.Collection(collectionName).InsertOne(context.Background(), order)
	if err != nil {
		return "", err
	}

	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("result id is not the ObjectID type")
	}

	return id.Hex(), nil
}

func (repository *OrderRepository) FindAll() ([]*model.Order, error) {
	cursor, err := repository.store.Collection(collectionName).Find(context.Background(), bson.D{})
	if err != nil {
		return []*model.Order{}, err
	}

	defer cursor.Close(context.Background())

	results := make([]*model.Order, 0)
	for cursor.Next(context.Background()) {
		var elem model.Order
		err := cursor.Decode(&elem)
		if err != nil {
			return []*model.Order{}, err
		}

		results = append(results, &elem)
	}

	if err := cursor.Err(); err != nil {
		return []*model.Order{}, err
	}
	return results, nil
}
