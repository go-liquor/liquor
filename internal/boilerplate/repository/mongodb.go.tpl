package db

import (
    "go.mongodb.org/mongo-driver/v2/mongo"
    "{{.module}}/internal/ports"
)

type {{.repositoryName}}Repository struct {
    db *mongo.Database
}

// New{{.repositoryName}}Repository create instance to ports.{{.repositoryName}}Repository
func New{{.repositoryName}}Repository(db *mongo.Database) ports.{{.repositoryName}}Repository {
    return &{{.repositoryName}}Repository{
        db: db,
    }
}