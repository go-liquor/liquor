package main

import (
	"github.com/go-liquor/liquor/v2/app"
	{{- if ne .Database "none" }}
	{{- if ne .Database "mongodb" }}
	"{{.Package}}/migrations"
	{{- end }}
	"{{.Package}}/app/adapters/database"
	{{- end }}
	
	"{{.Package}}/app/services"
)

func main() {
	app.New(
		{{- if ne .Database "none" }}
		migrations.Migrations,
		app.WithRepository(database.New{{.PascalCaseName}}Database),
		{{- end }}
		app.WithService(services.New{{.PascalCaseName}}Service),
		
	)
}
