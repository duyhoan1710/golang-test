package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	controllerInterface "api-orders/internal/interface/controller"
	serviceInterface "api-orders/internal/interface/service"
)

type userController struct {
	UserService serviceInterface.IUserService
}

func NewUserController(userService serviceInterface.IUserService) controllerInterface.IUserController {
	return &userController{
		UserService: userService,
	}
}

func (userController *userController) FindProfileById(c *gin.Context) {
	userId := c.GetString("x-user-id")

	profile, customError := userController.UserService.FindProfileById(c, userId)
	if customError != nil {
		c.JSON(customError.GetStatusCode(), customError)
		return
	}

	c.JSON(http.StatusOK, profile)
}
