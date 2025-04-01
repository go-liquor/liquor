package liquordb

import (
	"github.com/go-liquor/liquor/v2/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.uber.org/zap"
)

func NewMongoDBConnection(cfg *config.Config, logger *zap.Logger) (*mongo.Database, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(cfg.GetString(config.MongoDBDns)))
	if err != nil {
		logger.Fatal("Failed to connect to mongodb database", zap.Error(err))
		return nil, err
	}
	db := client.Database(cfg.GetString(config.MongoDBName))
	return db, nil
}
