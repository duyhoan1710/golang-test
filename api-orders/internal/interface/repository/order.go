package repository

import (
	model "api-orders/internal/model"
	"context"
)

type IOrderRepository interface {
	Create(c context.Context, order *model.Order) error
	UpdateOne(c context.Context, orderId string, order *model.Order) error
	Find(c context.Context) ([]model.Order, error)
	FindById(c context.Context, id string) (model.Order, error)
	FindByUserId(c context.Context, userId string) ([]model.Order, error)
}
