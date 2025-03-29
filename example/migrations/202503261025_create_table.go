package migrations

import (
	"context"
	"fmt"

	"github.com/go-liquor/liquor/v2/example/app/entity"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

func CreateTable(m *migrate.Migrations) {
	m.MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Println("[up] migrations")
		_, err := db.NewCreateTable().IfNotExists().Model((*entity.User)(nil)).Exec(ctx)
		return err
	}, func(ctx context.Context, db *bun.DB) error {
		fmt.Println("[down] migrations")
		_, err := db.NewDropTable().IfExists().Model((*entity.User)(nil)).Exec(ctx)
		return err
	})
}
