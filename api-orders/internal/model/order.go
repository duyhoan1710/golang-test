package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionOrder = "orders"
)

type Order struct {
	Id     primitive.ObjectID `bson:"_id"`
	State  int                `bson:"state"`
	UserId string             `bson:"userId"`
}
