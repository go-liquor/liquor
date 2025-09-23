package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
	"go.uber.org/fx"
)

// ConfigModule enable the config (is required)
var ConfigModule = fx.Module("liquor-config", fx.Provide(
	readConfigFile,
))

func readConfigFile() (*Config, error) {
	vp := viper.New()
	files := []string{
		"/etc/secrets/config.yaml",
		"config.yaml",
	}
	var file string
	for _, f := range files {
		if _, err := os.Stat(f); !os.IsNotExist(err) {
			file = f
			break
		}
	}
	if file == "" {
		return nil, fmt.Errorf("config file not found try: [%v]", files)
	}

	vp.AutomaticEnv()
	vp.SetConfigFile(file)
	vp.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := vp.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	return &Config{stg: vp}, nil
}
