package route

import (
	"api-orders/internal/bootstrap"

	"github.com/gin-gonic/gin"

	_ "api-orders/internal/dto"
)

// @Router			/orders [post]
// @Security BearerAuth
// @Summary      Create order router
// @Description  Create new order
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Param  payload  body  dto.OrderRequest  true  "Create new order"
// @Success      201  {object}  dto.OrderResponse
func CreateOrderRouter(app *bootstrap.Application, group *gin.RouterGroup) {
	orderController := app.Controllers.OrderController

	group.POST("/orders", orderController.CreateOrder)
}

// @Router			/orders/{orderId}/cancel [put]
// @Security BearerAuth
// @Summary      Cancel order router
// @Description  Cancel order
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Param orderId path string true "Order Id"
// @Success      204  {object}  dto.OrderResponse
func CancelOrderRouter(app *bootstrap.Application, group *gin.RouterGroup) {
	orderController := app.Controllers.OrderController

	group.PUT("/orders/:orderId/cancel", orderController.CancelOrder)
}

// @Router			/orders/{orderId} [get]
// @Security BearerAuth
// @Summary      Get detail order router
// @Description  Get detail order
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Param orderId path string true "Order Id"
// @Success      200  {object}  dto.OrderResponse
func FindOrderRouter(app *bootstrap.Application, group *gin.RouterGroup) {
	orderController := app.Controllers.OrderController

	group.GET("/orders/:orderId", orderController.FindById)
}

// @Router			/orders [get]
// @Security BearerAuth
// @Summary      Get list orders router
// @Description  Get list orders of user
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.OrderResponse
func ListOrdersRouter(app *bootstrap.Application, group *gin.RouterGroup) {
	orderController := app.Controllers.OrderController

	group.GET("/orders", orderController.ListOrders)
}
