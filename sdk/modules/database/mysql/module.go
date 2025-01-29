package mysql

import "go.uber.org/fx"

var DatabaseMysqlModule = fx.Module("liquor-database-mysql", fx.Provide(
	NewConnection,
))
