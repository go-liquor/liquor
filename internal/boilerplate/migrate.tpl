package migrations

import (
	"context"
	"fmt"
	{{- if .NewProject }}
	"{{.Package}}/app/domain/entity"
	{{- end }}
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

func {{.MigrateName}}(m *migrate.Migrations) {
	m.MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Println("[up] migrations")
		{{- if .NewProject }}
		_, err := db.NewCreateTable().IfNotExists().Model((*entity.{{.SingularPascalCaseName}})(nil)).Exec(ctx)
		return err
		{{- else }}
		return nil
		{{- end }}
	}, func(ctx context.Context, db *bun.DB) error {
		fmt.Println("[down] migrations")
		{{- if .NewProject }}
		_, err := db.NewDropTable().IfExists().Model((*entity.{{.SingularPascalCaseName}})(nil)).Exec(ctx)
		return err
		{{- else }}
		return nil
		{{- end }}

	})
}
