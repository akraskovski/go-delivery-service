package mongo

import (
	"github.com/akraskovski/go-delivery-service/internal/app/store"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store struct {
	db              *mongo.Client
	orderRepository *OrderRepository
}

func New(db *mongo.Client) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Order() store.OrderRepository {
	if s.orderRepository == nil {
		s.orderRepository = &OrderRepository{store: s}
	}

	return s.orderRepository
}
