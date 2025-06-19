package liquordb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"reflect"
)

type MongoFindBase struct {
	collectionName string
	err            error
	collection     any
	db             *mongo.Database
	filter         bson.M
	limit          int64
	skip           int64
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

		findOptions := options.Find()
		if m.limit != 0 {
			findOptions.SetLimit(m.limit)
		}
		if m.skip != 0 {
			findOptions.SetSkip(m.skip)
		}

		cursor, err := m.db.Collection(m.collectionName).Find(ctx, m.filter, findOptions)
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

func (m *MongoFindBase) Limit(limit int64) FindBase {
	m.limit = limit
	return m
}

func (m *MongoFindBase) Skip(skip int64) FindBase {
	m.skip = skip
	return m
}
