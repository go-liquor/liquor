package rest

import "go.uber.org/fx"

var RestModule = fx.Module("lq-adapter-rest", fx.Provide(
	instanceServer,
	newRoute,
), fx.Invoke(
	initialRoute,
	startServer,
))
