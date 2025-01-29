package sqlite

import "go.uber.org/fx"

var DatabaseSqliteModule = fx.Module("liquor-database-sqlite", fx.Provide(
	NewConnection,
))
