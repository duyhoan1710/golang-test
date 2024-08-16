package bootstrap

import (
	"api-orders/internal/interface/controller"
	"api-orders/internal/interface/repository"
	"api-orders/internal/interface/service"

	controllerImpl "api-orders/internal/api/controller"
	serviceImpl "api-orders/internal/api/service"
	repositoryImpl "api-orders/internal/repository"

	model "api-orders/internal/model"
	mongo "api-orders/internal/mongo"
)

type Repositories struct {
	UserRepository  repository.IUserRepository
	OrderRepository repository.IOrderRepository
}

type Services struct {
	AuthService  service.IAuthService
	UserService  service.IUserService
	OrderService service.IOrderService
}

type Controllers struct {
	AuthController  controller.IAuthController
	UserController  controller.IUserController
	OrderController controller.IOrderController
}

func InitDependency(db mongo.Database, app *Application) {
	userRepository := &repositoryImpl.UserRepository{Database: db, Collection: model.CollectionUser}
	orderRepository := &repositoryImpl.OrderRepository{Database: db, Collection: model.CollectionOrder}

	repositories := Repositories{
		UserRepository:  userRepository,
		OrderRepository: orderRepository,
	}

	authService := &serviceImpl.AuthService{
		UserRepository: userRepository,
		Env:            app.Env,
	}
	userService := &serviceImpl.UserService{
		UserRepository: userRepository,
	}
	orderService := &serviceImpl.OrderService{
		UserService:     userService,
		OrderRepository: orderRepository,
		Env:             app.Env,
	}

	services := Services{
		AuthService:  authService,
		UserService:  userService,
		OrderService: orderService,
	}

	controllers := Controllers{
		AuthController: &controllerImpl.AuthController{
			AuthService: authService,
		},
		UserController: &controllerImpl.UserController{
			UserService: userService,
		},
		OrderController: &controllerImpl.OrderController{
			OrderService: orderService,
		},
	}

	app.Repositories = repositories
	app.Services = services
	app.Controllers = controllers

}
