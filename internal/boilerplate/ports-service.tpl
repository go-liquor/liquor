package ports

import (
	"context"
	
	{{- if .NewProject }}
	"{{.Package}}/app/domain/entity"
	{{- end }}
)

type {{.PascalCaseName}}Service interface {
	{{- if .NewProject }}
	Get(ctx context.Context) []entity.{{.SingularPascalCaseName}}
	{{- end }}
}
