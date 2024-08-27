package service

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"api-orders/internal/enum"
	"api-orders/internal/exception"
	grpcClient "api-orders/internal/grpc/grpc-internal"
	"api-orders/internal/model"

	repositoryInterface "api-orders/internal/interface/repository"
	serviceInterface "api-orders/internal/interface/service"
)

type orderService struct {
	UserService       serviceInterface.IUserService
	OrderRepository   repositoryInterface.IOrderRepository
	PaymentGRPCClient grpcClient.IPaymentGRPCClient
}

func NewOrderService(userService serviceInterface.IUserService, orderRepository repositoryInterface.IOrderRepository, paymentGRPCClient grpcClient.IPaymentGRPCClient) serviceInterface.IOrderService {
	return &orderService{
		UserService:       userService,
		OrderRepository:   orderRepository,
		PaymentGRPCClient: paymentGRPCClient,
	}
}

func (orderService *orderService) CreateOrder(c context.Context, userId string) (customError exception.ICustomError) {
	orders, err := orderService.OrderRepository.FindByUserId(c, userId)
	if err != nil {
		return exception.NewCustomError(exception.INTERNAL_SERVER_ERROR, err.Error())
	}

	for _, order := range orders {
		if order.State == int(enum.Created) || order.State == int(enum.Confirmed) {
			return exception.NewCustomError(exception.LAST_ORDER_NOT_FINISH)
		}
	}

	order := model.Order{
		Id:     primitive.NewObjectID(),
		State:  int(enum.Created),
		UserId: userId,
	}

	err = orderService.OrderRepository.Create(c, &order)
	if err != nil {
		return exception.NewCustomError(exception.INTERNAL_SERVER_ERROR, err.Error())
	}

	isSuccess := orderService.PaymentGRPCClient.ProcessPayment(order.Id.Hex(), order.UserId)

	if isSuccess {
		order.State = int(enum.Confirmed)
		orderService.OrderRepository.UpdateOne(c, order.Id.Hex(), &order)

		time.AfterFunc(30*time.Second, func() {
			order.State = int(enum.Delivered)
			orderService.OrderRepository.UpdateOne(c, order.Id.Hex(), &order)
		})
	} else {
		order.State = int(enum.Cancelled)
		orderService.OrderRepository.UpdateOne(c, order.Id.Hex(), &order)
	}

	return customError
}

func (orderService *orderService) CancelOrder(c context.Context, userId string, orderId string) (customError exception.ICustomError) {
	var err error = nil

	order, err := orderService.OrderRepository.FindById(c, orderId)
	if err != nil {
		return exception.NewCustomError(exception.INTERNAL_SERVER_ERROR, err.Error())
	}

	if order.State == int(enum.Cancelled) || order.State == int(enum.Delivered) {
		return exception.NewCustomError(exception.CANNOT_CANNEL_ORDER)
	}

	order.State = int(enum.Cancelled)

	err = orderService.OrderRepository.UpdateOne(c, orderId, &order)
	if err != nil {
		return exception.NewCustomError(exception.INTERNAL_SERVER_ERROR, err.Error())
	}

	return customError
}

func (orderService *orderService) FindOrderById(c context.Context, userId string, orderId string) (order model.Order, customError exception.ICustomError) {
	order, err := orderService.OrderRepository.FindById(c, orderId)
	if err != nil {
		return order, exception.NewCustomError(exception.INTERNAL_SERVER_ERROR, err.Error())
	}

	return order, customError
}

func (orderService *orderService) ListOrders(c context.Context, userId string) (orders []model.Order, customError exception.ICustomError) {
	orders, err := orderService.OrderRepository.FindByUserId(c, userId)
	if err != nil {
		return orders, exception.NewCustomError(exception.INTERNAL_SERVER_ERROR, err.Error())
	}

	return orders, customError
}
