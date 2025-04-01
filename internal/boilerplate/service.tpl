package services

import (
	"context"

	{{- if .NewProject }}
	"{{.Package}}/app/domain/entity"
	{{- end }}
	"{{.Package}}/app/ports"
	"go.uber.org/zap"
)

type {{.PascalCaseName}}Service struct {
	logger *zap.Logger
	{{- if ne .Database "none" }}
	{{- if .NewProject }}
	repo   ports.{{.PascalCaseName}}Repository
	{{- end }}
	{{- end }}
}

func New{{.PascalCaseName}}Service(lg *zap.Logger,
{{- if ne .Database "none" }}
{{- if .NewProject }}
repo ports.{{.PascalCaseName}}Repository,
{{- end }}
{{- end }}
) ports.{{.PascalCaseName}}Service {
	return &{{.PascalCaseName}}Service{
		logger: lg,
        {{- if ne .Database "none" }}
		{{- if .NewProject }}
		repo:   repo,
		{{- end }}
        {{- end }}
	}
}

{{- if .NewProject }}
func (u *{{.PascalCaseName}}Service) Get(ctx context.Context) []entity.{{.SingularPascalCaseName}} {
	return make([]entity.{{.SingularPascalCaseName}}, 0)
}
{{- end }}
