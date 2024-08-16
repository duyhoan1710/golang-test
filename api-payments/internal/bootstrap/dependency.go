package bootstrap

import (
	mongo "api-payments/internal/mongo"
)

type Repositories struct {
}

type Services struct {
}

type Controllers struct {
}

func InitDependency(db mongo.Database, app *Application) {
	repositories := Repositories{}

	services := Services{}

	controllers := Controllers{}

	app.Repositories = repositories
	app.Services = services
	app.Controllers = controllers

}
