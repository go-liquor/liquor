package mongodb

import "go.uber.org/fx"

var DatabaseMongoDBModule = fx.Module("liquor-database-mongodb", fx.Provide(
	NewConnection,
	UseDatabase,
))
