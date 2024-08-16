package service

import (
	model "api-orders/internal/model"

	"github.com/gin-gonic/gin"
)

type IOrderService interface {
	CreateOrder(c *gin.Context, userId string)
	ListOrders(c *gin.Context, userId string) (orders []model.Order)
	FindOrderById(c *gin.Context, userId string, orderId string) (order model.Order)
	ChangeOrderToDelivered(c *gin.Context, order *model.Order) error
	CancelOrder(c *gin.Context, userId string, orderId string)
}
