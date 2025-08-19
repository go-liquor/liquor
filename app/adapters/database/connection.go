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

type driver string

const (
	MongoDB  driver = "mongodb"
	ORM      driver = "orm"
	MySQL    driver = "mysql"
	Postgres driver = "postgres"
	SQLITE   driver = "sqlite"
)

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

	for driverValue, _ := range cfg.GetStringMap(config.Database) {
		if driverValue == string(MongoDB) {
			logger.Info("creating connection with mongodb")
			client, err := mongo.Connect(options.Client().SetConnectTimeout(time.Second * 30).ApplyURI(cfg.GetString(config.DatabaseMongoDBUri)))
			if err != nil {
				logger.Fatal("Failed to connect to mongodb database", zap.Error(err))
				return result, err
			}
			db := client.Database(cfg.GetString(config.DatabaseMongoDBDBName))
			result.MongoDB = db
		}
		if driverValue == string(ORM) {
			switch cfg.GetString(config.DatabaseORMDriver) {
			case string(SQLITE):
				sqldb, err = sql.Open(sqliteshim.ShimName, cfg.GetString(config.DatabaseORMDNS))
			case string(MySQL):
				sqldb, err = sql.Open("mysql", cfg.GetString(config.DatabaseORMDNS))
			case string(Postgres):
				sqldb = sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(cfg.GetString(config.DatabaseORMDNS))))
			}
			if err != nil {
				return result, err
			}
			db := bun.NewDB(sqldb, sqlitedialect.New())
			result.ORM = db
		}
	}
	return result, nil
}
