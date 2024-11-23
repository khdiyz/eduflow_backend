package service

import (
	"eduflow/config"
	"eduflow/internal/models"
	"eduflow/internal/repository"
	"eduflow/internal/storage"
	"eduflow/pkg/logger"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	Role
	Authorization
	School
}

func NewServices(repo *repository.Repository, storage *storage.Storage, cfg *config.Config, logger *logger.Logger) *Service {
	return &Service{
		Role:          NewRoleService(repo),
		Authorization: NewAuthService(repo, cfg),
		School:        NewSchoolService(repo),
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

type School interface {
	Create(input models.CreateSchool) (uuid.UUID, error)
	GetListSchool(filter models.SchoolFilter) ([]models.School, int, error)
}
