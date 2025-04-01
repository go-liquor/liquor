package database

import (
	"database/sql"
	"github.com/go-liquor/liquor/v2/app/adapters/database/liquordb"
	"github.com/go-liquor/liquor/v2/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/driver/sqliteshim"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"log"
)

type ConnectionOutput struct {
	fx.Out
	DB  *bun.DB
	ODM liquordb.ODM
}

// NewConnection creates a new database connection using the Bun ORM.
// It supports multiple database drivers (SQLite, MySQL, PostgreSQL) based on configuration.
//
// Parameters:
//   - cfg: configuration object containing database settings
//   - logger: zap logger instance for error reporting
//
// Returns:
//   - *bun.DB: configured database connection instance
//
// The function will fatal log if connection fails.
// Supported drivers:
//   - SQLite
//   - MySQL
//   - PostgreSQL
func NewConnection(cfg *config.Config, logger *zap.Logger) ConnectionOutput {
	var sqldb *sql.DB
	var err error

	if cfg.GetString(config.DatabaseDriver) == config.MongoDBDriver {
		db, err := liquordb.NewMongoDBConnection(cfg, logger)
		if err != nil {
			log.Fatalf("Error connecting to database: %v", err)
		}
		res := liquordb.NewODMMongoDB(db)
		return ConnectionOutput{ODM: res}
	}

	switch cfg.GetString(config.DatabaseDriver) {
	case config.SQliteDriver:
		sqldb, err = sql.Open(sqliteshim.ShimName, cfg.GetString(config.SQliteDns))
	case config.MySQLDriver:
		sqldb, err = sql.Open("mysql", cfg.GetString(config.MySQLDns))
	case config.PostgresDriver:
		sqldb = sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(cfg.GetString(config.PostgresDns))))
	}
	if err != nil {
		logger.Fatal("failed to connect in database", zap.Error(err))
	}
	db := bun.NewDB(sqldb, sqlitedialect.New())
	return ConnectionOutput{DB: db}
}
