package mongo

import (
	"github.com/akraskovski/go-delivery-service/internal/app/model"
)

type OrderRepository struct {
	store *Store
}

func (repository *OrderRepository) Create(order *model.Order) (*model.Order, error) {
	return nil, nil
}

func (repository *OrderRepository) FindAll() ([]*model.Order, error) {
	return []*model.Order{}, nil
}
