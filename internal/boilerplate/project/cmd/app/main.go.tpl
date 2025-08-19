package main

import (
	"github.com/go-liquor/liquor/v3/app"
	"{{.module}}/internal/adapters/rest"
)

func main() {
	app.New(
		app.WithRestApi(rest.NewRestApi),
	)
}
