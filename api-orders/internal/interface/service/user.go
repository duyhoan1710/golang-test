package service

import (
	model "api-orders/internal/model"

	"github.com/gin-gonic/gin"
)

type IUserService interface {
	FindProfileById(c *gin.Context, userId string) (user model.User, isExist bool)
}
