package repository

import (
	"context"

	"api-orders/internal/mongo"

	"api-orders/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	repositoryInterface "api-orders/internal/interface/repository"
)

type userRepository struct {
	Database   mongo.Database
	Collection string
}

func NewUserRepository(database mongo.Database, collection string) repositoryInterface.IUserRepository {
	return &userRepository{
		Database:   database,
		Collection: collection,
	}
}

func (ur *userRepository) Create(c context.Context, user *model.User) error {
	collection := ur.Database.Collection(ur.Collection)

	_, err := collection.InsertOne(c, user)

	return err
}

func (ur *userRepository) Find(c context.Context) ([]model.User, error) {
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

func (ur *userRepository) FindByEmail(c context.Context, email string) (model.User, error) {
	collection := ur.Database.Collection(ur.Collection)
	var user model.User
	err := collection.FindOne(c, bson.M{"email": email}).Decode(&user)
	return user, err
}

func (ur *userRepository) FindById(c context.Context, id string) (model.User, error) {
	collection := ur.Database.Collection(ur.Collection)

	var user model.User

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&user)
	return user, err
}
