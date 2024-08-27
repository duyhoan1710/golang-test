package bootstrap

import (
	"api-orders/config"
	"api-orders/internal/mongo"
)

type Application struct {
	Env          *config.Env
	Mongo        mongo.Client
	Repositories Repositories
	Services     Services
	Controllers  Controllers
}

func App() Application {
	app := &Application{}
	app.Env = config.NewEnv()
	app.Mongo = NewMongoDatabase(app.Env)

	db := app.Mongo.Database(app.Env.DBName)
	InitDependency(db, app)

	return *app
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(app.Mongo)
}
