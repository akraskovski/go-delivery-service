package model

import "time"

type Order struct {
	Id             string
	Name           string
	DeliverAddress string
	DeliverTime    time.Time
}
