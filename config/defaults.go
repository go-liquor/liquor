package config

const (
	AppName               = "app.name"
	AppDebug              = "app.debug"
	GrpcPort              = "grpc.port"
	RestDisabled          = "rest.disabled"
	RestPort              = "rest.port"
	CorsDefault           = "rest.cors.default"
	CorsAllowOrigins      = "rest.cors.origins"
	CorsAllowMethods      = "rest.cors.methods"
	CorsAllowHeaders      = "rest.cors.headers"
	CorsAllowCredentials  = "rest.cors.credentials"
	LogLevel              = "log.level"
	LogFormat             = "log.format"
	Database              = "database"
	DatabaseMongoDBUri    = "database.mongodb.uri"
	DatabaseMongoDBDBName = "database.mongodb.dbName"
	DatabaseORMDriver     = "database.orm.driver"
	DatabaseORMDNS        = "database.orm.dns"
)
