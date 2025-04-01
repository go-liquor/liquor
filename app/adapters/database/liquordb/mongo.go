package liquordb

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoDBODM struct {
	db *mongo.Database
}

func (m *MongoDBODM) NewInsert(col any) InsertBase {
	collectionName, err := GetCollectionName(col)
	return &MongoInsertBase{
		collectionName: collectionName,
		err:            err,
		collection:     col,
		db:             m.db,
	}
}

func (m *MongoDBODM) NewFind(col any) FindBase {
	collectionName, err := GetCollectionName(col)
	return &MongoFindBase{
		collectionName: collectionName,
		err:            err,
		collection:     col,
		db:             m.db,
	}
}

func (m *MongoDBODM) NewDelete(col any) DeleteBase {
	collectionName, err := GetCollectionName(col)
	return &MongoDeleteBase{
		collectionName: collectionName,
		err:            err,
		collection:     col,
		db:             m.db,
	}
}

func (m *MongoDBODM) NewUpdate(col any) UpdateBase {
	collectionName, err := GetCollectionName(col)
	return &MongoUpdateBase{
		collectionName: collectionName,
		err:            err,
		collection:     col,
		db:             m.db,
	}
}

func (m *MongoDBODM) GetInstance() *mongo.Database {
	return m.db
}
