package store

type Store interface {
	Order() OrderRepository
}
