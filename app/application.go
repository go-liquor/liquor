package app

import (
	"context"
	"fmt"

	"github.com/go-liquor/liquor/v2/app/adapters/database"
	"github.com/go-liquor/liquor/v2/app/adapters/rest"
	"github.com/go-liquor/liquor/v2/config"
	"github.com/go-liquor/liquor/v2/logger"
	"github.com/go-liquor/liquor/v2/pkg/lqstring"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const (
	groupTagApiRest      = "liquor-restApis"
	groupTagRepositories = "liquor-repositories"
)

type Option = fx.Option

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
	return fx.Module("liquor-repository-"+unique, fx.Provide(
		repo))
}

func WithMigrations(migrations ...any) Option {
	migrations = append(migrations, func(db *bun.DB, migrations *migrate.Migrations, logger *zap.Logger) error {
		migrator := migrate.NewMigrator(db,
			migrations,
			migrate.WithTableName(database.MigrationsTableName),
			migrate.WithLocksTableName(database.MigrationsLocksTableName))
		if err := database.Init(context.TODO(), db); err != nil {
			return fmt.Errorf("failed to init migrate:%w", err)
		}
		group, err := migrator.Migrate(context.TODO())
		if err != nil {
			return fmt.Errorf("failed to migrate: %w", err)
		}
		if group.ID == 0 {
			logger.Info("there are no new migrations to run")
			return nil
		}
		logger.Info(fmt.Sprintf("migrated to %s", group))
		return nil
	})
	return fx.Module("liquor-migrations",
		fx.Provide(func() *migrate.Migrations {
			return migrate.NewMigrations()
		}),
		fx.Invoke(
			migrations...,
		))
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
	for _, option := range options {
		opts = append(opts, option)
	}

	app := fx.New(
		opts...,
	)
	app.Run()
}
