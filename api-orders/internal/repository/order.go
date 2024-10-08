package repository

import (
	"context"

	"api-orders/internal/mongo"

	"api-orders/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	repositoryInterface "api-orders/internal/interface/repository"
)

type orderRepository struct {
	Database   mongo.Database
	Collection string
}

func NewOrderRepository(database mongo.Database, collection string) repositoryInterface.IOrderRepository {
	return &orderRepository{
		Database:   database,
		Collection: collection,
	}
}

func (ur *orderRepository) Create(c context.Context, order *model.Order) error {
	collection := ur.Database.Collection(ur.Collection)

	_, err := collection.InsertOne(c, order)

	return err
}

func (ur *orderRepository) UpdateOne(c context.Context, orderId string, order *model.Order) error {
	collection := ur.Database.Collection(ur.Collection)

	idHex, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		return err
	}

	_, err = collection.UpdateOne(c, bson.M{"_id": idHex}, bson.M{
		"$set": order,
	})

	return err
}

func (ur *orderRepository) Find(c context.Context) ([]model.Order, error) {
	collection := ur.Database.Collection(ur.Collection)

	cursor, err := collection.Find(c, bson.D{})

	if err != nil {
		return nil, err
	}

	var orders []model.Order

	err = cursor.All(c, &orders)
	if orders == nil {
		return []model.Order{}, err
	}

	return orders, err
}

func (ur *orderRepository) FindById(c context.Context, id string) (model.Order, error) {
	collection := ur.Database.Collection(ur.Collection)

	var order model.Order

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return order, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&order)
	return order, err
}

func (ur *orderRepository) FindByUserId(c context.Context, userId string) ([]model.Order, error) {
	collection := ur.Database.Collection(ur.Collection)

	cursor, err := collection.Find(c, bson.M{"userId": userId})

	if err != nil {
		return nil, err
	}

	var orders []model.Order

	err = cursor.All(c, &orders)
	if orders == nil {
		return []model.Order{}, err
	}

	return orders, err
}
