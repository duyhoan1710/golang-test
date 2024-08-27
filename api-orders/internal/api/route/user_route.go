package route

import (
	"api-orders/internal/bootstrap"

	"github.com/gin-gonic/gin"
)

// @Router			 /users/profile [get]
// @Security BearerAuth
// @Summary      Get profile detail router
// @Description  Get profile detail by ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success      201  {object}  dto.ProfileResponse
func FindProfileRouter(app *bootstrap.Application, group *gin.RouterGroup) {
	userController := app.Controllers.UserController

	group.GET("/users/profile", userController.FindProfileById)
}
