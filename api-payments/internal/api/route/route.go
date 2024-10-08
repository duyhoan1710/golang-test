package route

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"api-payments/internal/api/middleware"
	"api-payments/internal/bootstrap"
)

func Setup(app *bootstrap.Application, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs
	// Swagger Endpoint
	publicRouter.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(app.Env.AccessTokenSecret))
	// All Private APIs

}
