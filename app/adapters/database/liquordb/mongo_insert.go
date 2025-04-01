package liquordb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"time"
)

type MongoInsertBase struct {
	collectionName string
	err            error
	collection     any
	db             *mongo.Database
}

func (i *MongoInsertBase) Exec(ctx context.Context) (any, error) {
	if i.err != nil {
		return "", i.err
	}
	content, err := bson.Marshal(i.collection)
	if err != nil {
		return "", fmt.Errorf("failed to marshal collection content: %w", err)
	}
	data := bson.M{}
	if err := bson.Unmarshal(content, &data); err != nil {
		return "", fmt.Errorf("failed to unmarshal collection content: %w", err)
	}
	data["createdAt"] = time.Now()
	data["updatedAt"] = time.Now()
	res, err := i.db.Collection(i.collectionName).InsertOne(ctx, data)
	if err != nil {
		return "", err
	}
	return res.InsertedID, err
}
