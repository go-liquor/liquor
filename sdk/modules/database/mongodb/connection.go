package mongodb

import (
	"github.com/go-liquor/liquor/sdk/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.uber.org/zap"
)

// NewConnection create a new connection to mongodb database
//
// Returns:
// - *mongo.Client: a new connection to mongodb database
func NewConnection(config *config.Config, logger *zap.Logger) *mongo.Client {
	client, err := mongo.Connect(options.Client().
		ApplyURI(config.GetString("database.mongodb.dsn")))
	if err != nil {
		logger.Fatal("failed to connect in database", zap.Error(err))
	}
	return client
}

// UseDatabase use a database in mongodb
//
// Returns:
// - *mongo.Database: a database in mongodb
func UseDatabase(client *mongo.Client, config *config.Config) *mongo.Database {
	return client.Database(config.GetString("database.mongodb.database"))
}
