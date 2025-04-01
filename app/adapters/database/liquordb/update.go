package liquordb

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type UpdateBase interface {
	Where(filter bson.M) UpdateBase
	Set(name string, value interface{}) UpdateBase
	Exec(ctx context.Context) error
}
