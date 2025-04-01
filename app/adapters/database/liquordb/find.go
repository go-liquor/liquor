package liquordb

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type FindBase interface {
	Where(filter bson.M) FindBase
	Scan(ctx context.Context) error
}
