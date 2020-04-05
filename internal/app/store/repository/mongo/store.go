package mongo

import (
	"github.com/akraskovski/go-delivery-service/internal/app/store"
)

type Store struct {
	db              *interface{} //todo replace by mongo driver
	orderRepository *OrderRepository
}

func New(db *interface{}) *Store {
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
