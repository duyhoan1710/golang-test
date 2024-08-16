package controller

import "github.com/gin-gonic/gin"

type IUserController interface {
	FindProfileById(c *gin.Context)
}
