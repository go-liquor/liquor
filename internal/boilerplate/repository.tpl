package database

import (
	"context"

	"{{.Package}}/app/entity"
	"{{.Package}}/app/ports"
	{{- if .Database "mongodb" }}
	"github.com/go-liquor/liquor/v2/app/adapters/database/liquordb"
	{{- else }}
	"github.com/uptrace/bun"
	{{- end }}
)

type {{.PascalCaseName}}Database struct {
    {{- if .Database "mongodb" }}
    db liquordb.ODM
    {{- else }}
	db *bun.DB
	{{- end }}
}
{{- if .Database "mongodb" }}
func New{{.PascalCaseName}}Database(db liquordb.ODM) ports.{{.PascalCaseName}}Repository {
{{- else }}
func New{{.PascalCaseName}}Database(db *bun.DB) ports.{{.PascalCaseName}}Repository {
{{- end }}
	return &{{.PascalCaseName}}Database{
		db: db,
	}
}

func (u *{{.PascalCaseName}}Database) Create(ctx context.Context, d *entity.{{.SingularPascalCaseName}}) error {
	_, err := u.db.NewInsert(d).Exec(ctx)
	return err
}

func (u *{{.PascalCaseName}}Database) Get(ctx context.Context) []entity.{{.SingularPascalCaseName}} {
	var {{.PascalCaseName}} []entity.{{.SingularPascalCaseName}}
	u.db.NewSelect({{.PascalCaseName}}).Scan(ctx)
	return {{.PascalCaseName}}
}
