package liquordb

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoDeleteBase struct {
	collectionName string
	err            error
	collection     any
	db             *mongo.Database
	filter         bson.M
}

func (m *MongoDeleteBase) Where(filter bson.M) DeleteBase {
	for name, value := range filter {
		filter[name] = value
	}
	return m
}

func (m *MongoDeleteBase) Exec(ctx context.Context) error {
	if m.err != nil {
		return m.err
	}
	_, err := m.db.Collection(m.collectionName).DeleteMany(ctx, m.filter)
	return err
}
