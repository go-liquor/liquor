package database

import (
	"context"

	"example/app/entity"
	"example/app/ports"
	"github.com/uptrace/bun"
)

type UsersDatabase struct {
	db *bun.DB
}

func NewUsersDatabase(db *bun.DB) ports.UserRepository {
	return &UsersDatabase{
		db: db,
	}
}

func (u *UsersDatabase) Create(ctx context.Context, user *entity.User) error {
	_, err := u.db.NewInsert().Model(user).Exec(ctx)
	return err
}

func (u *UsersDatabase) Get(ctx context.Context) []entity.User {
	var users []entity.User
	u.db.NewSelect().Model(&users).Scan(ctx)
	return users
}
