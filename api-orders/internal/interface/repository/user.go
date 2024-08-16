package repository

import (
	model "api-orders/internal/model"
	"context"
)

type IUserRepository interface {
	Create(c context.Context, user *model.User) error
	Find(c context.Context) ([]model.User, error)
	FindByEmail(c context.Context, email string) (model.User, error)
	FindById(c context.Context, id string) (model.User, error)
}
