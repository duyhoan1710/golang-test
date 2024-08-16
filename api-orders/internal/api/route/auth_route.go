package route

import (
	"api-orders/internal/bootstrap"

	"github.com/gin-gonic/gin"

	_ "api-orders/internal/dto"
)

// @Router			/login [post]
// @Summary      Login router
// @Description  Login into system
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param  payload  body  dto.LoginRequest  true  "Login into system"
// @Success      201  {object}  dto.LoginResponse
func LoginRouter(app *bootstrap.Application, group *gin.RouterGroup) {
	authController := app.Controllers.AuthController

	group.POST("/login", authController.Login)
}

// @Router			/signup [post]
// @Summary      Signup router
// @Description  Signup new account
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param  payload  body  dto.SignupRequest  true  "Signup new account"
// @Success      201  {object}  dto.SignupResponse
func SignupRouter(app *bootstrap.Application, group *gin.RouterGroup) {
	authController := app.Controllers.AuthController

	group.POST("/signup", authController.Signup)
}

// @Router			/refresh-token [post]
// @Summary      Refresh token router
// @Description  Refresh new token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param  payload  body  dto.RefreshTokenRequest  true  "Refresh new token"
// @Success      201  {object}  dto.RefreshTokenResponse
func RefreshTokenRouter(app *bootstrap.Application, group *gin.RouterGroup) {
	authController := app.Controllers.AuthController

	group.POST("/refresh-token", authController.RefreshToken)
}
