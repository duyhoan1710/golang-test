package controller

import (
	"api-orders/internal/api/service"
	"api-orders/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	OrderService *service.OrderService
}

func (orderController *OrderController) CreateOrder(c *gin.Context) {
	var request dto.OrderRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}

	userId := c.GetString("x-user-id")

	orderController.OrderService.CreateOrder(c, userId)

	c.JSON(http.StatusOK, nil)
}

func (orderController *OrderController) FindById(c *gin.Context) {
	orderId := c.Param("orderId")

	userId := c.GetString("x-user-id")

	order := orderController.OrderService.FindOrderById(c, userId, orderId)

	c.JSON(http.StatusOK, order)
}

func (orderController *OrderController) ListOrders(c *gin.Context) {
	userId := c.GetString("x-user-id")

	orders := orderController.OrderService.ListOrders(c, userId)

	c.JSON(http.StatusOK, orders)
}

func (orderController *OrderController) CancelOrder(c *gin.Context) {
	orderId := c.Param("orderId")
	userId := c.GetString("x-user-id")

	orderController.OrderService.CancelOrder(c, userId, orderId)

	c.JSON(http.StatusOK, nil)
}
