package http

import "go.uber.org/fx"

// httpModule enable the HTTP server
var httpModule = fx.Module("liquor-http", fx.Provide(
	instanceServer,
),
	fx.Invoke(
		startServer,
		initialRoute,
	))
