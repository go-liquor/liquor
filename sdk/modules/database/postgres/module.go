package postgres

import "go.uber.org/fx"

var DatabasePostgresModule = fx.Module("liquor-database-postgres", fx.Provide(
	NewConnection,
))
