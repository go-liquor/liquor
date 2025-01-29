package config

import "github.com/spf13/viper"

type Config struct {
	stg *viper.Viper
}

// Get retrieves a configuration value by its key as an interface{}.
//
// Parameters:
// - key: The key identifying the configuration value.
//
// Returns:
// - The value associated with the key as an interface{}.
func (c *Config) Get(key string) interface{} {
	return c.stg.Get(key)
}

// GetString retrieves a configuration value by its key as a string.
//
// Parameters:
// - key: The key identifying the configuration value.
//
// Returns:
// - The value associated with the key as a string.
func (c *Config) GetString(key string) string {
	return c.stg.GetString(key)
}

// GetStringSlice retrieves a configuration value by its key as a list of string.
//
// Parameters:
// - key: The key identifying the configuration value.
//
// Returns:
// - The value associated with the key as a list of string.
func (c *Config) GetStringSlice(key string) []string {
	return c.stg.GetStringSlice(key)
}

// GetInt retrieves a configuration value by its key as an int.
//
// Parameters:
// - key: The key identifying the configuration value.
//
// Returns:
// - The value associated with the key as an int.
func (c *Config) GetInt(key string) int {
	return c.stg.GetInt(key)
}

// GetInt64 retrieves a configuration value by its key as an int64.
//
// Parameters:
// - key: The key identifying the configuration value.
//
// Returns:
// - The value associated with the key as an int64.
func (c *Config) GetInt64(key string) int64 {
	return c.stg.GetInt64(key)
}

// GetBool retrieves a configuration value by its key as a bool.
//
// Parameters:
// - key: The key identifying the configuration value.
//
// Returns:
// - The value associated with the key as a bool.
func (c *Config) GetBool(key string) bool {
	return c.stg.GetBool(key)
}

// GetFloat64 retrieves a configuration value by its key as a float64.
//
// Parameters:
// - key: The key identifying the configuration value.
//
// Returns:
// - The value associated with the key as a float64.
func (c *Config) GetFloat64(key string) float64 {
	return c.stg.GetFloat64(key)
}

// GetAppName retrieves the name of the application from the configuration.
//
// Returns:
// - The application name as a string.
func (c *Config) GetAppName() string {
	return c.GetString("app.name")
}

// IsDebug checks whether the application is in debug mode.
//
// Returns:
// - true if the application is in debug mode, false otherwise.
func (c *Config) IsDebug() bool {
	return c.GetBool("app.debug")
}

// GetServerHttpPort retrieves the HTTP server port from the configuration.
//
// Returns:
// - The HTTP server port as an int64.
func (c *Config) GetServerHttpPort() int64 {
	return c.GetInt64("server.http.port")
}

// GetServerHttpCorsDefaultAllow checks if CORS is enabled by default for the HTTP server.
//
// Returns:
// - true if CORS is enabled by default, false otherwise.
func (c *Config) GetServerHttpCorsDefaultAllow() bool {
	return c.GetBool("server.http.cors.default")
}

// GetServerHttpCorsAllowOrigins retrieves the list of allowed origins for CORS on the HTTP server.
//
// Returns:
// - A slice of strings containing the allowed origins.
func (c *Config) GetServerHttpCorsAllowOrigins() []string {
	return c.stg.GetStringSlice("server.http.cors.origins")
}

// GetServerHttpCorsAllowMethods retrieves the list of allowed HTTP methods for CORS on the HTTP server.
//
// Returns:
// - A slice of strings containing the allowed HTTP methods.
func (c *Config) GetServerHttpCorsAllowMethods() []string {
	return c.stg.GetStringSlice("server.http.cors.methods")
}

// GetServerHttpCorsAllowHeaders retrieves the list of allowed headers for CORS on the HTTP server.
//
// Returns:
// - A slice of strings containing the allowed headers.
func (c *Config) GetServerHttpCorsAllowHeaders() []string {
	return c.stg.GetStringSlice("server.http.cors.headers")
}

// GetServerHttpCorsAllowCredentials checks if credentials are allowed in CORS requests for the HTTP server.
//
// Returns:
// - true if credentials are allowed, false otherwise.
func (c *Config) GetServerHttpCorsAllowCredentials() bool {
	return c.GetBool("server.http.cors.credentials")
}

// GetServerGrpcPort retrieves the gRPC server port from the configuration.
//
// Returns:
// - The gRPC server port as an int64.
func (c *Config) GetServerGrpcPort() int64 {
	return c.GetInt64("server.grpc.port")
}

// GetPassswordBcryptCost retrieves the bcrypt cost for password hashing.
//
// Returns:
// - The bcrypt cost as an int.
func (c *Config) GetPassswordBcryptCost() int {
	return c.GetInt("password.bcrypt.cost")
}

// GetDatabaseDriver retrieves the database driver name from the configuration.
//
// Returns:
// - The database driver name as a string.
func (c *Config) GetDatabaseDriver() string {
	return c.GetString("database.driver")
}

// GetLogLevel retrieves the log level from the configuration.
//
// Returns:
// - The log level as a string (can be info,warn,error,debug).
func (c *Config) GetLogLevel() string {
	return c.GetString("log.level")
}

// GetLogFormat retrieves the log format from the configuration.
//
// Returns:
// - The log format as a string (can be json, console).
func (c *Config) GetLogFormat() string {
	return c.GetString("log.format")
}
