package service_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"api-orders/internal/exception"
	repositoryMock "api-orders/internal/repository/mocks"

	"api-orders/internal/api/service"

	"api-orders/internal/model"
)

func TestFindProfileById(t *testing.T) {
	userObjectId := primitive.NewObjectID()
	userId := userObjectId.String()

	mockProfile := model.User{
		ID:    userObjectId,
		Name:  "test",
		Email: "test@gmail.com",
	}

	t.Run("Should return error", func(t *testing.T) {
		mockUserRepository := new(repositoryMock.UserRepositoryMock)

		userService := service.NewUserService(
			mockUserRepository,
		)

		mockUserRepository.On("FindById", mock.Anything, userId).Return(mockProfile, errors.New(""))

		_, err := userService.FindProfileById(context.Background(), userId)

		if err.GetErrorCode() != exception.INTERNAL_SERVER_ERROR.Index() {
			t.Fatalf("Error code should be INTERNAL SERVER ERROR")
		}
	})

	t.Run("Should success", func(t *testing.T) {
		mockUserRepository := new(repositoryMock.UserRepositoryMock)

		userService := service.NewUserService(
			mockUserRepository,
		)

		mockUserRepository.On("FindById", mock.Anything, userId).Return(mockProfile, nil)

		_, err := userService.FindProfileById(context.Background(), userId)

		fmt.Printf("%v", err)

		if err != nil {
			t.Fatalf("Should return success")
		}
	})
}
