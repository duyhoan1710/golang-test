package service

import (
	"api-orders/internal/exception"
	"api-orders/internal/model"
	"context"
)

type IOrderService interface {
	CreateOrder(c context.Context, userId string) exception.ICustomError
	ListOrders(c context.Context, userId string) ([]model.Order, exception.ICustomError)
	FindOrderById(c context.Context, userId string, orderId string) (model.Order, exception.ICustomError)
	CancelOrder(c context.Context, userId string, orderId string) exception.ICustomError
}
