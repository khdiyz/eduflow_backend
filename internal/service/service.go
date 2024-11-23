package service

import (
	"eduflow/config"
	"eduflow/internal/models"
	"eduflow/internal/repository"
	"eduflow/internal/storage"
	"eduflow/pkg/logger"
	"time"
)

type Service struct {
	Role
	Authorization
}

func NewServices(repo *repository.Repository, storage *storage.Storage, cfg *config.Config, logger *logger.Logger) *Service {
	return &Service{
		Role:          NewRoleService(repo),
		Authorization: NewAuthService(repo, cfg),
	}
}

type Role interface {
	Create(role models.CreateRole) error
}

type Authorization interface {
	CreateToken(user models.User, tokenType string, expiresAt time.Time) (*models.Token, error)
	GenerateTokens(user models.User) (*models.Token, *models.Token, error)
	ParseToken(token string) (*jwtCustomClaim, error)
	Login(request models.LoginRequest) (*models.Token, *models.Token, error)
}
