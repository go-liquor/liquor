package main

import (
	"example/app/adapters/database"
	"example/app/adapters/rest"
	"example/app/services"
	"example/migrations"
	"github.com/go-liquor/liquor/v2/app"
)

func main() {
	app.New(
		app.WithService(services.NewUserService),
		app.WithRepository(database.NewUsersDatabase),
		app.WithRestApi(rest.NewUsersApi),
		migrations.Migrations,
	)
}
