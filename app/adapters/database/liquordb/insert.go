package liquordb

import "context"

type InsertBase interface {
	Exec(ctx context.Context) (any, error)
}
