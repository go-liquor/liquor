package database

import "go.uber.org/fx"

var Module = fx.Module("liquor-adapter-database", fx.Provide(
	NewConnection,
	NewProvider,
))
