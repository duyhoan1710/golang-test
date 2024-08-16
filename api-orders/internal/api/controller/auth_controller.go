package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"api-orders/internal/api/service"
	"api-orders/internal/dto"
)

type AuthController struct {
	AuthService *service.AuthService
}

func (authController *AuthController) Login(c *gin.Context) {
	var request dto.LoginRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}

	accessToken, refreshToken := authController.AuthService.Login(c, request.Email, request.Password)

	loginResponse := dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, loginResponse)
}

func (authController *AuthController) Signup(c *gin.Context) {
	var request dto.SignupRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}

	accessToken, refreshToken := authController.AuthService.Signup(c, request.Name, request.Email, request.Password)

	signupResponse := dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, signupResponse)
}

func (authController *AuthController) RefreshToken(c *gin.Context) {
	var request dto.RefreshTokenRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}

	accessToken, refreshToken := authController.AuthService.RefreshToken(c, request.RefreshToken)

	refreshTokenResponse := dto.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, refreshTokenResponse)
}
