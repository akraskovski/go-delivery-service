package model

import "time"

type Order struct {
	Id             string
	name           string
	deliverAddress string
	deliverTime    time.Time
}
