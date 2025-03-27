package logger

import (
	"go.uber.org/fx"
)

// LoggerModule enable the logger (is required)
var LoggerModule = fx.Module("liquor-logger", fx.Provide(
	instanceLogger,
))
