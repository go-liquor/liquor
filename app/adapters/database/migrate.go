package database

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

const (
	MigrationsTableName      = "liquor_migrations"
	MigrationsLocksTableName = "liquor_migrations_locks"
)

type Migration struct {
	ID         int64 `bun:",pk,type:serial"`
	Name       string
	GroupID    int64
	MigratedAt time.Time `bun:",notnull,nullzero,default:current_timestamp"`
}

type migrationLock struct {
	ID        int64  `bun:",pk,type:serial"`
	TableName string `bun:",unique"`
}

func Init(ctx context.Context, db *bun.DB) error {
	if _, err := db.NewCreateTable().
		Model((*Migration)(nil)).
		ModelTableExpr(MigrationsTableName).
		IfNotExists().
		Exec(ctx); err != nil {
		return err
	}
	if _, err := db.NewCreateTable().
		Model((*migrationLock)(nil)).
		ModelTableExpr(MigrationsLocksTableName).
		IfNotExists().
		Exec(ctx); err != nil {
		return err
	}
	return nil
}
