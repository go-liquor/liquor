package database

import "go.uber.org/fx"

var Module = fx.Module("lq-adapter-database", fx.Provide(
	NewConnection,
))
