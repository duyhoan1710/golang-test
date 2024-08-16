package service

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"api-orders/internal/dto"
	model "api-orders/internal/model"
	"api-orders/internal/repository"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func (userService *UserService) FindProfileById(c *gin.Context, userId string) (user model.User, isExist bool) {
	user, err := userService.UserRepository.FindById(c, userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	return user, true
}
