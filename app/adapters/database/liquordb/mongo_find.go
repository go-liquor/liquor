package liquordb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"reflect"
)

type MongoFindBase struct {
	collectionName string
	err            error
	collection     any
	db             *mongo.Database
	filter         bson.M
}

func (m *MongoFindBase) Where(filter bson.M) FindBase {
	if m.filter == nil {
		m.filter = bson.M{}
	}
	for name, value := range filter {
		m.filter[name] = value
	}
	return m
}

func (m *MongoFindBase) Scan(ctx context.Context) error {
	if m.err != nil {
		return m.err
	}
	t := reflect.TypeOf(m.collection)
	// If collection is a *pointer FindOne, else FindMany
	if t.Kind() == reflect.Ptr {
		// is the only obj
		if t.Elem().Name() != "" {
			return m.db.Collection(m.collectionName).FindOne(ctx, m.filter).Decode(m.collection)
		}
		cursor, err := m.db.Collection(m.collectionName).Find(ctx, m.filter)
		if err != nil {
			return err
		}
		return cursor.All(ctx, m.collection)
	}

	return fmt.Errorf("failed to read as pointer or slice")
}

func (m *MongoFindBase) Count(ctx context.Context) (int64, error) {
	return m.db.Collection(m.collectionName).CountDocuments(ctx, m.filter)
}
