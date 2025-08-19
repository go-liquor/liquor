package rest

import "github.com/go-liquor/liquor/v3/app/adapters/rest"

type RestApi struct {
}

func NewRestApi() rest.Api {
	return &RestApi{}
}

func (s *RestApi) Routes(r *rest.Route) {

}
