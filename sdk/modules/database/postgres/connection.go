package postgres

import (
	"database/sql"

	"github.com/go-liquor/liquor/sdk/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"go.uber.org/zap"
)

// NewConnection create a new connection to postgres database
//
// Returns:
// - *bun.DB: a new connection to postgres database
func NewConnection(config *config.Config, logger *zap.Logger) *bun.DB {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(config.GetString("database.sqlite.dsn"))))

	db := bun.NewDB(sqldb, pgdialect.New())
	return db
}
