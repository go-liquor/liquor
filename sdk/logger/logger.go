package logger

import (
	"log"

	"github.com/go-liquor/liquor/sdk/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func instanceLogger(config *config.Config) *zap.Logger {
	cfg := zap.NewProductionConfig()
	cfg.DisableCaller = true
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.Encoding = "json"

	if config.GetLogFormat() == "console" {
		cfg.Encoding = "console"
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	var level zapcore.Level
	switch viper.GetString(config.GetLogLevel()) {
	case "debug":
		cfg.DisableCaller = false
		cfg.DisableStacktrace = false
		level = zapcore.DebugLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "dpanic":
		level = zapcore.DPanicLevel
	case "panic":
		level = zapcore.PanicLevel
	case "fatal":
		level = zapcore.FatalLevel
	default:
		level = zapcore.InfoLevel
	}
	cfg.Level = zap.NewAtomicLevelAt(level)

	logger, err := cfg.Build()
	if err != nil {
		log.Fatalf("failed to build logger: %v", err)
	}
	return logger
}
