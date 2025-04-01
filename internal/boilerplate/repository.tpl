package database

import (
	"context"

	"{{.Package}}/app/domain/entity"
	"{{.Package}}/app/ports"
	{{- if eq .Database "mongodb" }}
	"github.com/go-liquor/liquor/v2/app/adapters/database/liquordb"
	{{- else }}
	"github.com/uptrace/bun"
	{{- end }}
)

type {{.PascalCaseName}}Database struct {
    {{- if eq .Database "mongodb" }}
    db liquordb.ODM
    {{- else }}
	db *bun.DB
	{{- end }}
}
{{- if eq .Database "mongodb" }}
func New{{.PascalCaseName}}Database(db liquordb.ODM) ports.{{.PascalCaseName}}Repository {
{{- else }}
func New{{.PascalCaseName}}Database(db *bun.DB) ports.{{.PascalCaseName}}Repository {
{{- end }}
	return &{{.PascalCaseName}}Database{
		db: db,
	}
}

func (u *{{.PascalCaseName}}Database) Create(ctx context.Context, d *entity.{{.SingularPascalCaseName}}) error {
    {{- if eq .Database "mongodb" }}
	_, err := u.db.NewInsert(d).Exec(ctx)
	{{- else }}
	_, err := u.db.NewInsert().Model(d).Exec(ctx)
	{{- end }}
	return err
}

func (u *{{.PascalCaseName}}Database) Get(ctx context.Context) []entity.{{.SingularPascalCaseName}} {
	var {{.PascalCaseName}} []entity.{{.SingularPascalCaseName}}
	{{- if eq .Database "mongodb" }}
	u.db.NewFind({{.PascalCaseName}}).Scan(ctx)
	{{- else }}
	u.db.NewSelect().Model(&{{.PascalCaseName}}).Scan(ctx)
	{{- end }}
	return {{.PascalCaseName}}
}
