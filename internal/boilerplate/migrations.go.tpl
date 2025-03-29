package migrations

import "github.com/go-liquor/liquor/v2/app"

var Migrations = app.WithMigrations(
    {{.MigrateName}},
    // TODO: add migrations
)
