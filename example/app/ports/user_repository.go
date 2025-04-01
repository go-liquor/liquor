package ports

import (
	"context"

	"example/app/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	Get(ctx context.Context) []entity.User
}
