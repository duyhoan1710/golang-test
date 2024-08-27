package service

import (
	"api-orders/internal/exception"
	"api-orders/internal/model"
	"context"
)

type IUserService interface {
	FindProfileById(c context.Context, userId string) (user model.User, err exception.ICustomError)
}
