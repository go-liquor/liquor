package db

import (
    "github.com/uptrace/bun"
    "{{.module}}/internal/ports"
)

type {{.repositoryName}}Repository struct {
    db *bun.DB
}

// New{{.repositoryName}}Repository create instance to ports.{{.repositoryName}}Repository
func New{{.repositoryName}}Repository(db *bun.DB) ports.{{.repositoryName}}Repository {
    return &{{.repositoryName}}Repository{
        db: db,
    }
}