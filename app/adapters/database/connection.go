package database

import (
	"database/sql"
	"time"

	"github.com/go-liquor/liquor/v3/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/driver/sqliteshim"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const (
	MongoDB  string = "mongodb"
	ORM      string = "orm"
	MySQL    string = "mysql"
	Postgres string = "postgres"
	SQLITE   string = "sqlite"
)

var Drivers map[string]struct{} = map[string]struct{}{
	MongoDB: {},
	ORM:     {},
}

type ConnectionOutput struct {
	fx.Out
	ORM     *bun.DB
	MongoDB *mongo.Database
}

// NewConnection creates a new database connection using the Bun ORM OR Mongodb.
func NewConnection(cfg *config.Config, logger *zap.Logger) (ConnectionOutput, error) {
	var sqldb *sql.DB
	var err error
	result := ConnectionOutput{}

	if value := cfg.GetString(config.DatabaseMongoDBUri); value != "" {
		logger.Info("creating connection with mongodb")
		client, err := mongo.Connect(options.Client().SetConnectTimeout(time.Second * 30).ApplyURI(cfg.GetString(config.DatabaseMongoDBUri)))
		if err != nil {
			logger.Fatal("Failed to connect to mongodb database", zap.Error(err))
			return result, err
		}
		db := client.Database(cfg.GetString(config.DatabaseMongoDBDBName))
		result.MongoDB = db
	}

	if value := cfg.GetString(config.DatabaseORMDriver); value != "" {
		switch cfg.GetString(config.DatabaseORMDriver) {
		case SQLITE:
			sqldb, err = sql.Open(sqliteshim.ShimName, cfg.GetString(config.DatabaseORMDNS))
		case MySQL:
			sqldb, err = sql.Open("mysql", cfg.GetString(config.DatabaseORMDNS))
		case Postgres:
			sqldb = sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(cfg.GetString(config.DatabaseORMDNS))))
		}
		if err != nil {
			return result, err
		}
		db := bun.NewDB(sqldb, sqlitedialect.New())
		result.ORM = db
	}

	return result, nil
}
