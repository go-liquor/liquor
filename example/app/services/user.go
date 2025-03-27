package services

import (
	"context"
	"fmt"

	"github.com/go-liquor/liquor/v2/example/app/entity"
	"github.com/go-liquor/liquor/v2/example/app/ports"
	"go.uber.org/zap"
)

type UserService struct {
	logger *zap.Logger
	repo   ports.UserRepository
}

func NewUserService(lg *zap.Logger, repo ports.UserRepository) ports.UserService {
	return &UserService{
		logger: lg,
		repo:   repo,
	}
}

func (u *UserService) Get(ctx context.Context) []entity.User {
	return u.repo.Get(ctx)
}

func (u *UserService) Create(ctx context.Context, user *entity.User) error {
	err := u.repo.Create(ctx, user)
	if err != nil {
		u.logger.Error("failed to create user", zap.Error(err))
		return fmt.Errorf("failed to create user")
	}
	u.logger.Debug("user created")
	return nil

}
