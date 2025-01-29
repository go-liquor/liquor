package sqlite

import (
	"database/sql"

	"github.com/go-liquor/liquor/sdk/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"go.uber.org/zap"
)

// NewConnection create a new connection to sqlite database
//
// Returns:
// - *bun.DB: a new connection to sqlite database
func NewConnection(config *config.Config, logger *zap.Logger) *bun.DB {
	sqldb, err := sql.Open(sqliteshim.ShimName, config.GetString("database.sqlite.dsn"))
	if err != nil {
		logger.Fatal("failed to connect in database", zap.Error(err))
	}

	db := bun.NewDB(sqldb, sqlitedialect.New())
	return db
}
