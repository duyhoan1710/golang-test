package route

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"api-orders/internal/api/middleware"
	"api-orders/internal/bootstrap"
)

func Setup(app *bootstrap.Application, gin *gin.Engine) {
	publicRouter := gin.Group("")
	protectedRouter := gin.Group("")

	publicRouter.Use(middleware.InternalServerErrorHandler)
	protectedRouter.Use(middleware.InternalServerErrorHandler)

	// All Public APIs
	LoginRouter(app, publicRouter)
	SignupRouter(app, publicRouter)
	RefreshTokenRouter(app, publicRouter)
	publicRouter.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(app.Env.AccessTokenSecret))

	// All Private APIs
	FindProfileRouter(app, protectedRouter)

	CreateOrderRouter(app, protectedRouter)
	FindOrderRouter(app, protectedRouter)
	ListOrdersRouter(app, protectedRouter)
	CancelOrderRouter(app, protectedRouter)
}
