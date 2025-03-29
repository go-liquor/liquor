package database

import (
	"context"

	"{{.Package}}/app/entity"
	"{{.Package}}/app/ports"
	"github.com/uptrace/bun"
)

type {{.PascalCaseName}}Database struct {
	db *bun.DB
}

func New{{.PascalCaseName}}Database(db *bun.DB) ports.{{.PascalCaseName}}Repository {
	return &{{.PascalCaseName}}Database{
		db: db,
	}
}

func (u *{{.PascalCaseName}}Database) Create(ctx context.Context, d *entity.{{.SingularPascalCaseName}}) error {
	_, err := u.db.NewInsert().Model(d).Exec(ctx)
	return err
}

func (u *{{.PascalCaseName}}Database) Get(ctx context.Context) []entity.{{.SingularPascalCaseName}} {
	var {{.PascalCaseName}} []entity.{{.SingularPascalCaseName}}
	u.db.NewSelect().Model(&{{.PascalCaseName}}).Scan(ctx)
	return {{.PascalCaseName}}
}
