package service

import (
	"context"

	"api-orders/internal/model"

	repositoryInterface "api-orders/internal/interface/repository"
	serviceInterface "api-orders/internal/interface/service"

	exception "api-orders/internal/exception"
)

type userService struct {
	UserRepository repositoryInterface.IUserRepository
}

func NewUserService(userRepository repositoryInterface.IUserRepository) serviceInterface.IUserService {
	return &userService{
		UserRepository: userRepository,
	}
}

func (userService *userService) FindProfileById(c context.Context, userId string) (user model.User, customError exception.ICustomError) {
	user, err := userService.UserRepository.FindById(c, userId)

	if err != nil {
		customError = exception.NewCustomError(exception.INTERNAL_SERVER_ERROR)
	}

	return user, customError
}
