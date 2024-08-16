package repository

import (
	"context"

	mongo "api-payments/internal/mongo"

	model "api-payments/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	Database   mongo.Database
	Collection string
}

type IUserRepository interface {
	Create(c context.Context, user *model.User) error
	Find(c context.Context) ([]model.User, error)
	FindByEmail(c context.Context, email string) (model.User, error)
	FindById(c context.Context, id string) (model.User, error)
}

func (ur *UserRepository) Create(c context.Context, user *model.User) error {
	collection := ur.Database.Collection(ur.Collection)

	_, err := collection.InsertOne(c, user)

	return err
}

func (ur *UserRepository) Find(c context.Context) ([]model.User, error) {
	collection := ur.Database.Collection(ur.Collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var users []model.User

	err = cursor.All(c, &users)
	if users == nil {
		return []model.User{}, err
	}

	return users, err
}

func (ur *UserRepository) FindByEmail(c context.Context, email string) (model.User, error) {
	collection := ur.Database.Collection(ur.Collection)
	var user model.User
	err := collection.FindOne(c, bson.M{"email": email}).Decode(&user)
	return user, err
}

func (ur *UserRepository) FindById(c context.Context, id string) (model.User, error) {
	collection := ur.Database.Collection(ur.Collection)

	var user model.User

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&user)
	return user, err
}
