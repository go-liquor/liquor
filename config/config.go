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

func (c *Config) GetStringMap(key string) map[string]interface{} {
	return c.stg.GetStringMap(key)
}
