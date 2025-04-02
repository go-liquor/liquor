package liquordb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"time"
)

type MongoUpdateBase struct {
	collectionName string
	err            error
	collection     any
	db             *mongo.Database
	filter         bson.M
	setters        bson.M
}

func (m *MongoUpdateBase) Where(filter bson.M) UpdateBase {
	if m.filter == nil {
		m.filter = bson.M{}
	}
	for name, value := range filter {
		filter[name] = value
	}
	return m
}

func (m *MongoUpdateBase) Set(name string, value interface{}) UpdateBase {
	if m.setters == nil {
		m.setters = bson.M{}
	}
	m.setters[name] = value
	return m
}

func (m *MongoUpdateBase) Exec(ctx context.Context) error {
	if m.err != nil {
		return m.err
	}
	setters := m.setters
	if setters == nil {
		content, err := bson.Marshal(m.collection)
		if err != nil {
			return fmt.Errorf("failed to marshal collection content: %w", err)
		}
		data := bson.M{}
		if err := bson.Unmarshal(content, &data); err != nil {
			return fmt.Errorf("failed to unmarshal collection content: %w", err)
		}
		setters = data
	}
	setters["updatedAt"] = time.Now()
	_, err := m.db.Collection(m.collectionName).UpdateMany(ctx, m.filter, bson.M{"$set": setters})
	return err
}
