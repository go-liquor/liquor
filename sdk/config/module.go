package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"go.uber.org/fx"
)

// ConfigModule enable the config (is required)
var ConfigModule = fx.Module("liquor-config", fx.Provide(
	readConfigFile,
))

func readConfigFile() *Config {
	vp := viper.New()
	vp.SetConfigFile("config.yaml")
	vp.AutomaticEnv()
	if err := vp.ReadInConfig(); err != nil {
		fmt.Printf("failed to read config file: %v\n", err)
		os.Exit(1)
	}
	return &Config{stg: vp}
}
