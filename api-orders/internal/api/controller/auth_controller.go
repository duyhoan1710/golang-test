package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"api-orders/internal/dto"
	"api-orders/internal/exception"

	controllerInterface "api-orders/internal/interface/controller"
	serviceInterface "api-orders/internal/interface/service"
)

type authController struct {
	AuthService serviceInterface.IAuthService
}

func NewAuthController(authService serviceInterface.IAuthService) controllerInterface.IAuthController {
	return &authController{
		AuthService: authService,
	}
}

func (authController *authController) Login(c *gin.Context) {
	var request dto.LoginRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, exception.NewCustomError(exception.VALIDATION_ERROR, err.Error()))
		return
	}

	accessToken, refreshToken, customError := authController.AuthService.Login(c, request.Email, request.Password)
	if customError != nil {
		c.JSON(customError.GetStatusCode(), customError)
		return
	}

	loginResponse := dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, loginResponse)
}

func (authController *authController) Signup(c *gin.Context) {
	var request dto.SignupRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, exception.NewCustomError(exception.VALIDATION_ERROR, err.Error()))
		return
	}

	accessToken, refreshToken, customError := authController.AuthService.Signup(c, request.Name, request.Email, request.Password)
	if customError != nil {
		c.JSON(customError.GetStatusCode(), customError)
		return
	}

	signupResponse := dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, signupResponse)
}

func (authController *authController) RefreshToken(c *gin.Context) {
	var request dto.RefreshTokenRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, exception.NewCustomError(exception.VALIDATION_ERROR, err.Error()))
		return
	}

	accessToken, refreshToken, customError := authController.AuthService.RefreshToken(c, request.RefreshToken)
	if customError != nil {
		c.JSON(customError.GetStatusCode(), customError)
		return
	}

	refreshTokenResponse := dto.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, refreshTokenResponse)
}
