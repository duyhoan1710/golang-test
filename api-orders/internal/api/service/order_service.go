package service

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"api-orders/config"
	"api-orders/internal/dto"
	enum "api-orders/internal/enum"
	grpc_internal "api-orders/internal/grpc-gateway/grpc-internal"
	model "api-orders/internal/model"
	"api-orders/internal/repository"
)

type OrderService struct {
	UserService     *UserService
	OrderRepository *repository.OrderRepository
	Env             *config.Env
}

func (orderService *OrderService) ChangeOrderToDelivered(c *gin.Context, order *model.Order) error {
	order.State = int(enum.Delivered)

	err := orderService.OrderRepository.UpdateOne(c, order.Id.String(), order)

	return err
}

func (orderService *OrderService) CreateOrder(c *gin.Context, userId string) {
	orders, err := orderService.OrderRepository.FindByUserId(c, userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	for _, order := range orders {
		if order.State == int(enum.Created) || order.State == int(enum.Confirmed) {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Finish previous order first"})
			return
		}
	}

	order := model.Order{
		Id:     primitive.NewObjectID(),
		State:  int(enum.Created),
		UserId: userId,
	}

	err = orderService.OrderRepository.Create(c, &order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	isSuccess := grpc_internal.ProcessPayment(orderService.Env, order.Id.String(), order.UserId)

	if isSuccess {
		order.State = int(enum.Confirmed)
		orderService.OrderRepository.UpdateOne(c, order.Id.String(), &order)

		time.AfterFunc(30*time.Second, func() {
			order.State = int(enum.Delivered)
			orderService.OrderRepository.UpdateOne(c, order.Id.String(), &order)
		})
	} else {
		order.State = int(enum.Cancelled)
		orderService.OrderRepository.UpdateOne(c, order.Id.String(), &order)
	}
}

func (orderService *OrderService) CancelOrder(c *gin.Context, userId string, orderId string) {
	var err error = nil

	order, err := orderService.OrderRepository.FindById(c, orderId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	if order.State == int(enum.Cancelled) || order.State == int(enum.Delivered) {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Cannot cancel this order"})
		return
	}

	order.State = int(enum.Cancelled)

	err = orderService.OrderRepository.UpdateOne(c, orderId, &order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}
}

func (orderService *OrderService) FindOrderById(c *gin.Context, userId string, orderId string) (order model.Order) {
	var err error = nil

	_, isExist := orderService.UserService.FindProfileById(c, userId)
	if !isExist {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Message: "User not found with the given id"})
		return
	}

	order, err = orderService.OrderRepository.FindById(c, orderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	return order
}

func (orderService *OrderService) ListOrders(c *gin.Context, userId string) (orders []model.Order) {
	var err error = nil

	_, isExist := orderService.UserService.FindProfileById(c, userId)
	if !isExist {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Message: "User not found with the given id"})
		return
	}

	orders, err = orderService.OrderRepository.FindByUserId(c, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	return orders
}
