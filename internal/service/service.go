package service

import (
	"eduflow/config"
	"eduflow/internal/repository"
	"eduflow/internal/storage"
	"eduflow/pkg/logger"
)

type Service struct {
}

func NewService(repos *repository.Repository, storage *storage.Storage, config *config.Config, loggers *logger.Logger) *Service {
	return &Service{}
}
