package mysql

import (
	"database/sql"

	"github.com/go-liquor/liquor/sdk/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"go.uber.org/zap"
)

// NewConnection create a new connection to mysql database
//
// Returns:
// - *bun.DB: a new connection to mysql database
func NewConnection(config *config.Config, logger *zap.Logger) *bun.DB {
	sqldb, err := sql.Open("mysql", config.GetString("database.mysql.dsn"))
	if err != nil {
		logger.Fatal("failed to connect in database", zap.Error(err))
	}

	db := bun.NewDB(sqldb, mysqldialect.New())
	return db
}
