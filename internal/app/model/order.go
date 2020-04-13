package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Order struct {
	Id             primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Name           string             `json:"name" bson:"name"`
	DeliverAddress string             `json:"deliverAddress" bson:"deliverAddress"`
	DeliverTime    time.Time          `json:"deliverTime" bson:"deliverTime"`
}
