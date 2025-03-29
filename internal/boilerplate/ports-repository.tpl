package ports

import (
	"context"

	"{{.Package}}/app/entity"
)

type {{.PascalCaseName}}Repository interface {
	Create(ctx context.Context, d *entity.{{.SingularPascalCaseName}}) error
	Get(ctx context.Context) []entity.{{.SingularPascalCaseName}}
}
