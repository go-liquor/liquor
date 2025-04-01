package liquordb

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
	"time"
)

type ODM interface {
	NewInsert(col any) InsertBase
	NewUpdate(col any) UpdateBase
	NewDelete(col any) DeleteBase
	NewFind(col any) FindBase
	GetInstance() *mongo.Database
}

type Collection struct {
	CreatedAt time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

func NewODMMongoDB(db *mongo.Database) ODM {
	return &MongoDBODM{
		db: db,
	}
}
