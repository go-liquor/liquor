package database

import (
	"context"

	"github.com/uptrace/bun"
)

type Provider struct {
	db *bun.DB
}

func NewProvider(db *bun.DB) *Provider {
	return &Provider{
		db: db,
	}
}

// CreateTableIfNotExists creates a new database table for the given model if it doesn't already exist.
// It uses the Bun ORM's schema creation capabilities to generate the table structure.
//
// Parameters:
//   - n: any struct that represents a database model
//
// Returns:
//   - error: returns nil if successful, otherwise returns the error that occurred
//
// Example:
//
//	type User struct {
//	    ID   int64  `bun:"id,pk,autoincrement"`
//	    Name string `bun:"name,notnull"`
//	}
//
//	err := provider.CreateTableIfNotExists(&User{})
func (p *Provider) CreateTableIfNotExists(n any) error {
	_, err := p.db.NewCreateTable().IfNotExists().Model(n).Exec(context.TODO())
	return err
}

// CreateIndexIfNotExists creates a new database index for the given model if it doesn't already exist.
// It uses the Bun ORM's schema creation capabilities to generate the index.
//
// Parameters:
//   - n: any struct that represents a database model
//   - name: the name of the index to be created
//   - colums: variadic list of column names to be included in the index
//
// Returns:
//   - error: returns nil if successful, otherwise returns the error that occurred
//
// Example:
//
//	type User struct {
//	    ID    int64  `bun:"id,pk,autoincrement"`
//	    Email string `bun:"email,notnull"`
//	}
//
//	err := provider.CreateIndexIfNotExists(&User{}, "idx_user_email", "email")
func (p *Provider) CreateIndexIfNotExists(n any, name string, colums ...string) error {
	_, err := p.db.NewCreateIndex().IfNotExists().Model(n).Index(name).Column(colums...).Exec(context.TODO())
	return err
}
