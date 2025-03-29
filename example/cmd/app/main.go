package main

import (
	"github.com/go-liquor/liquor/v2/app"
	"github.com/go-liquor/liquor/v2/example/app/adapters/database"
	"github.com/go-liquor/liquor/v2/example/app/adapters/rest"
	"github.com/go-liquor/liquor/v2/example/app/services"
	"github.com/go-liquor/liquor/v2/example/migrations"
)

func main() {
	app.New(
		app.WithService(services.NewUserService),
		app.WithRepository(database.NewUsersDatabase),
		app.WithRestApi(rest.NewUsersApi),
		migrations.Migrations,
	)
}
