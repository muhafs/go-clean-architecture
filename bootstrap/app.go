package bootstrap

import "gorm.io/gorm"

type Application struct {
	Env   *Env
	Mysql *gorm.DB
}

func App() Application {
	app := new(Application)
	app.Env = NewEnv()
	app.Mysql = NewMySQLDatabase(app.Env)

	return *app
}

func (app *Application) CloseDBConnection() {
	CloseMysqlDBConnection(app.Mysql)
}
