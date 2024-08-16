package service

import "github.com/gin-gonic/gin"

type IAuthService interface {
	Login(c *gin.Context, email string, password string) (accessToken string, refreshToken string)
	Signup(c *gin.Context, name string, email string, password string) (accessToken string, refreshToken string)
	RefreshToken(c *gin.Context, currentRefreshToken string) (accessToken string, refreshToken string)
}
