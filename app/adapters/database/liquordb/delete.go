package liquordb

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type DeleteBase interface {
	Where(filter bson.M) DeleteBase
	Exec(ctx context.Context) error
}
