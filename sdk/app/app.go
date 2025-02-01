package app

import (
	"github.com/go-liquor/liquor/sdk/config"
	"github.com/go-liquor/liquor/sdk/logger"
	"go.uber.org/fx"
)

// NewApp create a new app
func NewApp(modules ...fx.Option) {
	options := []fx.Option{
		config.ConfigModule,
		logger.LoggerModule,
	}
	options = append(options, modules...)
	app := fx.New(options...)
	app.Run()
}

// RegisterService register you services
func RegisterService(services ...interface{}) fx.Option {
	return fx.Module("liquor-app-services", fx.Provide(
		services...,
	))
}

// RegisterRepositories register you repositories
func RegisterRepositories(repos ...interface{}) fx.Option {
	return fx.Module("liquor-app-repositories", fx.Provide(
		repos...,
	))
}
