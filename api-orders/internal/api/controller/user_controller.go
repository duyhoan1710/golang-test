package controller

import (
	"net/http"

	"api-orders/internal/api/service"

	"github.com/gin-gonic/gin"
)

type IUserController interface {
	FindProfileById(c *gin.Context)
}

type UserController struct {
	UserService *service.UserService
}

func (userController *UserController) FindProfileById(c *gin.Context) {
	userId := c.GetString("x-user-id")

	profile, _ := userController.UserService.FindProfileById(c, userId)

	c.JSON(http.StatusOK, profile)
}
