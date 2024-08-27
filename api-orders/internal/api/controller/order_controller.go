package controller

import (
	"api-orders/internal/dto"
	"api-orders/internal/exception"
	"net/http"

	"github.com/gin-gonic/gin"

	controllerInterface "api-orders/internal/interface/controller"
	serviceInterface "api-orders/internal/interface/service"
)

type orderController struct {
	OrderService serviceInterface.IOrderService
}

func NewOrderController(orderService serviceInterface.IOrderService) controllerInterface.IOrderController {
	return &orderController{
		OrderService: orderService,
	}
}

func (orderController *orderController) CreateOrder(c *gin.Context) {
	var request dto.OrderRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, exception.NewCustomError(exception.VALIDATION_ERROR, err.Error()))
		return
	}

	userId := c.GetString("x-user-id")

	customError := orderController.OrderService.CreateOrder(c, userId)
	if customError != nil {
		c.JSON(customError.GetStatusCode(), customError)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (orderController *orderController) FindById(c *gin.Context) {
	orderId := c.Param("orderId")

	userId := c.GetString("x-user-id")

	order, customError := orderController.OrderService.FindOrderById(c, userId, orderId)
	if customError != nil {
		c.JSON(customError.GetStatusCode(), customError)
		return
	}

	c.JSON(http.StatusOK, order)
}

func (orderController *orderController) ListOrders(c *gin.Context) {
	userId := c.GetString("x-user-id")

	orders, customError := orderController.OrderService.ListOrders(c, userId)
	if customError != nil {
		c.JSON(customError.GetStatusCode(), customError)
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (orderController *orderController) CancelOrder(c *gin.Context) {
	orderId := c.Param("orderId")
	userId := c.GetString("x-user-id")

	customError := orderController.OrderService.CancelOrder(c, userId, orderId)
	if customError != nil {
		c.JSON(customError.GetStatusCode(), customError)
		return
	}

	c.JSON(http.StatusOK, nil)
}
