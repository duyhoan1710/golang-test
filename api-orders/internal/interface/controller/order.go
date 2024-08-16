package controller

import "github.com/gin-gonic/gin"

type IOrderController interface {
	CreateOrder(c *gin.Context)
	FindById(c *gin.Context)
	ListOrders(c *gin.Context)
	CancelOrder(c *gin.Context)
}
