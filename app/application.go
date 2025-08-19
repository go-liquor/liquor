package app

import (
	"os"

	"github.com/go-liquor/liquor/v3/app/adapters/database"
	"github.com/go-liquor/liquor/v3/app/adapters/rest"
	"github.com/go-liquor/liquor/v3/config"
	"github.com/go-liquor/liquor/v3/logger"
	"github.com/go-liquor/liquor/v3/pkg/lqstring"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	groupTagApiRest = "lq-restApis"
)

type Option = fx.Option

// Register register a new provider
var Register = fx.Provide

// Call invoke a function
var Call = fx.Invoke

// WithRestApi registers a REST API implementation with the application.
// It wraps the provided function as a REST API provider and adds it to the "restApis" group
// for automatic registration during application startup.
//
// Parameters:
//   - fn: any function that returns a REST API implementation
//
// Returns:
//   - Option: a configured fx.Option that can be passed to New()
//
// Example:
//
//	app.New(
//	    WithRestApi(NewUserApi),
//	    WithRestApi(NewProductApi),
//	)
func WithRestApi(fn any) Option {
	return fx.Provide(
		fx.Annotate(
			fn,
			fx.As(new(rest.Api)),
			fx.ResultTags(`group:"`+groupTagApiRest+`"`),
		),
	)
}

// WithService registers a service provider with the application.
// It provides a simple way to add service implementations to the dependency injection container.
//
// Parameters:
//   - svc: any function or struct that implements a service
//
// Returns:
//   - Option: a configured fx.Option that can be passed to New()
//
// Example:
//
//	app.New(
//	    WithService(NewUserService),
//	    WithService(NewAuthService),
//	)
func WithService(svc any) Option {
	return fx.Provide(
		svc,
	)
}

// WithRepository registers a database repository implementation with the application.
// It wraps the provided repository as a database.Repository and adds it to the "liquor-repositories" group
// for automatic table creation during application startup.
//
// Parameters:
//   - repo: any function or struct that implements database.Repository interface
//
// Returns:
//   - Option: a configured fx.Option that can be passed to New()
//
// Example:
//
//	app.New(
//	    WithRepository(NewUserRepository),
//	    WithRepository(NewProductRepository),
//	)
func WithRepository(repo any) Option {
	unique := lqstring.RandomString(6)
	return fx.Module("lq-repo-"+unique, fx.Provide(
		repo))
}

// nologger print fx events when debug is enabled
func nologger(cfg *config.Config) fxevent.Logger {
	if cfg.GetBool("app.debug") || os.Getenv("DEBUG") == "on" {
		zapCfg := zap.NewProductionConfig()
		zapCfg.EncoderConfig.TimeKey = ""
		zapCfg.EncoderConfig.LevelKey = ""
		zapCfg.EncoderConfig.MessageKey = "msg"
		zapCfg.EncoderConfig.CallerKey = ""     // remove caller
		zapCfg.EncoderConfig.StacktraceKey = "" // remove stacktrace
		zapCfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
		l, err := zapCfg.Build()
		if err != nil {
			panic(err)
		}

		return &fxevent.ZapLogger{Logger: l}
	}
	return fxevent.NopLogger
}

// New creates and runs a new application instance using uber-fx dependency injection.
// It initializes the application with default modules (config, logger, and REST)
// and accepts additional options for customization.
//
// Parameters:
//   - options: variadic Option parameters that allow extending the application
//     with additional modules, providers, or invocations.
//
// The application will:
//   - Set up configuration management
//   - Initialize logging
//   - Configure REST server
//   - Register all REST APIs
//   - Start all provided services
func New(options ...Option) {
	opts := []fx.Option{
		fx.WithLogger(nologger),
		config.ConfigModule,
		logger.LoggerModule,
		rest.RestModule,
		database.Module,
		fx.Invoke(
			func(l *zap.Logger, cfg *config.Config) {
				l.Debug("Starting application " + cfg.GetString(config.AppName))
			},
			// registers rest apis
			func(s *rest.Route, api []rest.Api) {
				for _, a := range api {
					a.Routes(s)
				}
			},
		),
		// create a slice of rest apis
		fx.Provide(
			fx.Annotate(func(routes []rest.Api) []rest.Api { return routes }, fx.ParamTags(`group:"`+groupTagApiRest+`"`)),
		),
	}
	opts = append(opts, options...)

	app := fx.New(
		opts...,
	)
	app.Run()
}
