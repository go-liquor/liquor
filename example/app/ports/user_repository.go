package ports

import (
	"context"

	"github.com/go-liquor/liquor/v2/example/app/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	Get(ctx context.Context) []entity.User
}
