package service

import (
	"eduflow/config"
	"eduflow/internal/models"
	"eduflow/internal/repository"
	"eduflow/internal/storage"
	"eduflow/pkg/logger"
)

type Service struct {
	Role
}

func NewServices(repo *repository.Repository, storage *storage.Storage, config *config.Config, logger *logger.Logger) *Service {
	return &Service{
		Role: NewRoleService(repo),
	}
}

type Role interface {
	Create(role models.CreateRole) error
}
