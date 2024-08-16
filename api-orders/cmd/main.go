package main

import (
	"github.com/gin-gonic/gin"

	_ "api-orders/docs"
	route "api-orders/internal/api/route"
	"api-orders/internal/bootstrap"
)

// @title Sotatek Test Api Orders
// @version 0.0.1
// @description This is a sample server test.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	app := bootstrap.App()

	env := app.Env

	defer app.CloseDBConnection()

	gin := gin.Default()

	route.Setup(&app, gin)

	gin.Run(env.ServerAddress)
}
