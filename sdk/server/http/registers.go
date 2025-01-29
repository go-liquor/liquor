package http

import "go.uber.org/fx"

// RegisterRoutes register you routes
func RegisterRoutes(routes ...interface{}) fx.Option {
	return fx.Invoke(routes...)
}

// RegisterHandlers register you handlers
func RegisterHandlers(handlers ...interface{}) fx.Option {
	return fx.Provide(handlers...)
}

type Handler []interface{}
type Route []interface{}

type HttpModuleParam struct {
	Handlers Handler
	Routes   Route
}

func StartModule(input *HttpModuleParam) fx.Option {
	return fx.Module("liquor-app-http-server",
		httpModule,
		RegisterHandlers(input.Handlers...),
		RegisterRoutes(input.Routes...))
}
