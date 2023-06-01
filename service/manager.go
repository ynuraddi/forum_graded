package service

import (
	"graded/config"
	"graded/logger"
	"graded/repository"
)

type Manager struct{}

func Init(config *config.Config, logger *logger.Logger, repository *repository.Manager) (*Manager, error) {
	return nil, nil
}
