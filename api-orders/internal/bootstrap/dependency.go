package bootstrap

import (
	grpc_internal "api-orders/internal/grpc/grpc-internal"
	"api-orders/internal/interface/controller"
	"api-orders/internal/interface/repository"
	"api-orders/internal/interface/service"

	controllerImpl "api-orders/internal/api/controller"
	serviceImpl "api-orders/internal/api/service"
	repositoryImpl "api-orders/internal/repository"

	"api-orders/internal/model"
	"api-orders/internal/mongo"
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
	paymentGRPCClient, err := grpc_internal.NewPaymentGRPCClient(app.Env.GRPC_PAYMENT_SERVER_ADDRESS)
	if err != nil {
		panic(err)
	}

	userRepository := repositoryImpl.NewUserRepository(db, model.CollectionUser)
	orderRepository := repositoryImpl.NewOrderRepository(db, model.CollectionOrder)

	repositories := Repositories{
		UserRepository:  userRepository,
		OrderRepository: orderRepository,
	}

	authService := serviceImpl.NewAuthService(
		userRepository,
		app.Env,
	)
	userService := serviceImpl.NewUserService(
		userRepository,
	)
	orderService := serviceImpl.NewOrderService(
		userService,
		orderRepository,
		paymentGRPCClient,
	)

	services := Services{
		AuthService:  authService,
		UserService:  userService,
		OrderService: orderService,
	}

	controllers := Controllers{
		AuthController: controllerImpl.NewAuthController(
			authService,
		),
		UserController: controllerImpl.NewUserController(
			userService,
		),
		OrderController: controllerImpl.NewOrderController(
			orderService,
		),
	}

	app.Repositories = repositories
	app.Services = services
	app.Controllers = controllers

}
