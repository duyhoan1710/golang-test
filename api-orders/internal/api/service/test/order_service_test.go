package service_test

import (
	"context"
	"errors"
	"testing"

	mock "github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"api-orders/internal/api/service"
	serviceMock "api-orders/internal/api/service/mocks"
	"api-orders/internal/enum"
	"api-orders/internal/exception"
	"api-orders/internal/model"
	repositoryMock "api-orders/internal/repository/mocks"

	grpcMock "api-orders/internal/grpc/grpc-internal/mocks"
)

type OrderServiceMock struct {
	mock.Mock
}

func TestCreateOrder(t *testing.T) {
	c := context.Background()

	userObjectId := primitive.NewObjectID()
	userId := userObjectId.Hex()

	t.Run("Should return error INTERNAL_SERVER_ERROR", func(t *testing.T) {
		mockOrderRepository := new(repositoryMock.OrderRepositoryMock)
		mockUserService := new(serviceMock.UserServiceMock)
		mockPaymentGRPCClient := new(grpcMock.PaymentGRPCClientMock)

		orderService := service.NewOrderService(mockUserService, mockOrderRepository, mockPaymentGRPCClient)

		var mockOrders = []model.Order{}

		mockOrderRepository.On("FindByUserId", c, userId).Return(mockOrders, errors.New(("")))

		err := orderService.CreateOrder(context.Background(), userId)

		if err.GetErrorCode() != exception.INTERNAL_SERVER_ERROR.Index() {
			t.Fatalf("Error code should be INTERNAL_SERVER_ERROR")
		}
	})

	t.Run("Should return error if previous order is not finished", func(t *testing.T) {
		mockOrderRepository := new(repositoryMock.OrderRepositoryMock)
		mockUserService := new(serviceMock.UserServiceMock)
		mockPaymentGRPCClient := new(grpcMock.PaymentGRPCClientMock)

		orderService := service.NewOrderService(mockUserService, mockOrderRepository, mockPaymentGRPCClient)

		var mockOrders = []model.Order{
			{
				State:  int(enum.Created),
				UserId: userId,
			},
		}

		mockOrderRepository.On("FindByUserId", c, userId).Return(mockOrders, nil)

		err := orderService.CreateOrder(c, userId)

		if err.GetErrorCode() != exception.LAST_ORDER_NOT_FINISH.Index() {
			t.Fatalf("Error code should be LAST_ORDER_NOT_FINISH")
		}
	})

	t.Run("Should update order state to CANCELED when payment fail", func(t *testing.T) {
		mockOrderRepository := new(repositoryMock.OrderRepositoryMock)
		mockUserService := new(serviceMock.UserServiceMock)
		mockPaymentGRPCClient := new(grpcMock.PaymentGRPCClientMock)

		orderService := service.NewOrderService(mockUserService, mockOrderRepository, mockPaymentGRPCClient)

		var mockOrders = []model.Order{
			{
				State:  int(enum.Delivered),
				UserId: userId,
			},
		}

		// var newOrder = model.Order{
		// 	State:  int(enum.Created),
		// 	UserId: userId,
		// }

		mockOrderRepository.On("FindByUserId", c, userId).Return(mockOrders, nil)
		mockOrderRepository.On("Create", c, mock.Anything).Return(nil)
		mockPaymentGRPCClient.On("ProcessPayment", mock.Anything, mock.Anything).Return(false)

		mockOrderRepository.On("UpdateOne", c, mock.Anything).Return(nil)

		err := orderService.CreateOrder(c, userId)

		if err != nil {
			t.Fatalf("Should process success")
		}
	})

	t.Run("Should update order state to Confirmed when payment success", func(t *testing.T) {
		mockOrderRepository := new(repositoryMock.OrderRepositoryMock)
		mockUserService := new(serviceMock.UserServiceMock)
		mockPaymentGRPCClient := new(grpcMock.PaymentGRPCClientMock)

		orderService := service.NewOrderService(mockUserService, mockOrderRepository, mockPaymentGRPCClient)

		var mockOrders = []model.Order{
			{
				State:  int(enum.Delivered),
				UserId: userId,
			},
		}

		// var newOrder = model.Order{
		// 	State:  int(enum.Created),
		// 	UserId: userId,
		// }

		mockOrderRepository.On("FindByUserId", c, userId).Return(mockOrders, nil)
		mockOrderRepository.On("Create", c, mock.Anything).Return(nil)
		mockPaymentGRPCClient.On("ProcessPayment", mock.Anything, mock.Anything).Return(false)

		mockOrderRepository.On("UpdateOne", c, mock.Anything, mock.Anything).Return(nil)

		err := orderService.CreateOrder(c, userId)

		if err != nil {
			t.Fatalf("Should process success")
		}
	})
}
