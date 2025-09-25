package database

import (
	"database/sql"
	"os"
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

const (
	EnvMongoDBEnabled  = "LIQUOR_DB_MONGODB_ENABLED"
	EnvMongoDBURI      = "LIQUOR_DB_MONGODB_URI"
	EnvMongoDBDatabase = "LIQUOR_DB_MONGODB_DATABASE"

	EnvORMEnabled = "LIQUOR_DB_ORM_ENABLED"
	EnvORMDriver  = "LIQUOR_DB_ORM_DRIVER"
	EnvORMDNS     = "LIQUOR_DB_ORM_DNS"
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
	result := ConnectionOutput{}

	if os.Getenv(EnvMongoDBEnabled) == "true" {
		logger.Info("creating connection with mongodb")
		client, err := mongo.Connect(options.Client().SetConnectTimeout(time.Second * 30).ApplyURI(os.Getenv(EnvMongoDBURI)))
		if err != nil {
			logger.Fatal("Failed to connect to mongodb database", zap.Error(err))
			return result, err
		}
		db := client.Database(os.Getenv(EnvMongoDBDatabase))
		result.MongoDB = db
	}

	if os.Getenv(EnvORMEnabled) == "true" {
		dns := os.Getenv(EnvORMDNS)
		var err error
		switch os.Getenv(EnvORMDriver) {
		case SQLITE:
			sqldb, err = sql.Open(sqliteshim.ShimName, dns)
		case MySQL:
			sqldb, err = sql.Open("mysql", dns)
		case Postgres:
			sqldb = sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dns)))
		}
		if err != nil {
			return result, err
		}
		db := bun.NewDB(sqldb, sqlitedialect.New())
		result.ORM = db
	}

	return result, nil
}
