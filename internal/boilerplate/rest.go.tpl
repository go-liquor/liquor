package rest

import "github.com/go-liquor/liquor/v3/app/adapters/rest"

type {{.restName}}Api struct {
}

func New{{.restName}}Api() rest.Api {
	return &{{.restName}}Api{}
}

func (s *{{.restName}}Api) Routes(r *rest.Route) {
    group := r.Group("{{.group}}")
    {
        group.Get("/", func(r *rest.Request){
            r.Status(http.StatusOK)
        })
    }
}
