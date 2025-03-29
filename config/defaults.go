package config

const (
	AppName              = "app.name"
	AppDebug             = "app.debug"
	RestPort             = "rest.port"
	CorsDefault          = "rest.cors.default"
	CorsAllowOrigins     = "rest.cors.origins"
	CorsAllowMethods     = "rest.cors.methods"
	CorsAllowHeaders     = "rest.cors.headers"
	CorsAllowCredentials = "rest.cors.credentials"
	LogLevel             = "log.level"
	LogFormat            = "log.format"
	DatabaseDriver       = "database.driver"
	SQliteDns            = "database.sqlite.dns"
	PostgresDns          = "database.postgres.dns"
	MySQLDns             = "database.mysql.dns"
)

type DBDriver string

const (
	SQliteDriver   = "sqlite"
	PostgresDriver = "postgres"
	MySQLDriver    = "mysql"
)
