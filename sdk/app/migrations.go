package app

import "go.uber.org/fx"

// RegisterMigrations register the functions with migrations function
func RegisterMigrations(migrations ...any) fx.Option {
	return fx.Module("liquor-app-migrations", fx.Invoke(
		migrations...,
	))
}
