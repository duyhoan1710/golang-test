package controller

import "github.com/gin-gonic/gin"

type IAuthController interface {
	Login(c *gin.Context)
	Signup(c *gin.Context)
	RefreshToken(c *gin.Context)
}
