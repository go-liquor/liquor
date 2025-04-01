package rest

import (
	"net/http"

	"github.com/go-liquor/liquor/v2/app/adapters/rest"
	{{- if .NewProject }}
	"{{.Package}}/app/ports"
	{{- end }}
)

type {{.PascalCaseName}}Api struct {
	{{- if .NewProject }}
	svc ports.{{.PascalCaseName}}Service
	{{- end }}
}

{{- if .NewProject }}
func New{{.PascalCaseName}}Api(svc ports.{{.PascalCaseName}}Service) rest.Api {
{{- else }}
func New{{.PascalCaseName}}Api() rest.Api {
{{- end }}
	return &{{.PascalCaseName}}Api{
		svc: svc,
	}
}

func (u *{{.PascalCaseName}}Api) Routes(s *rest.Route) {
	s.Get("/", u.Get)
}

func (u *{{.PascalCaseName}}Api) Get(r *rest.Request) {
	r.Status(http.StatusOK)
}


