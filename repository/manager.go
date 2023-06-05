package repository

import (
	"context"

	"graded/config"
	"graded/logger"
	"graded/model"
	"graded/repository/sqlite"
)

type IUserRepository interface {
	Create(ctx context.Context, user model.User) error
	GetByID(ctx context.Context, uid int64) (model.User, error)
	GetByEmail(ctx context.Context, email string) (model.User, error)
}

type Manager struct {
	User IUserRepository
}

func Init(cfg *config.Config, lg *logger.Logger) (*Manager, error) {
	_, err := sqlite.OpenDB(cfg)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
