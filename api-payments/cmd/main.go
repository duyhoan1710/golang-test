package main

import (
	"api-payments/internal/bootstrap"
	grpc_internal "api-payments/internal/grpc-gateway/grpc-internal"
)

// @title Sotatek Test Api Payments
// @version 0.0.1
// @description This is a sample server test.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	app := bootstrap.App()

	env := app.Env

	defer app.CloseDBConnection()

	// gin := gin.Default()

	// route.Setup(&app, gin)

	// gin.Run(env.ServerAddress)

	grpc_internal.StartGRPCServer(env)
}
