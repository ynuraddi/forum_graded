package service

import (
	"graded/config"
	"graded/logger"
	"graded/repository"
)

type Manager struct{}

func Init(cfg *config.Config, lg *logger.Logger, repo *repository.Manager) (*Manager, error) {
	return nil, nil
}
