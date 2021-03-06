package store

import "github.com/akraskovski/go-delivery-service/internal/app/model"

type OrderRepository interface {
	Create(order *model.Order) (string, error)
	FindAll() ([]*model.Order, error)
}
