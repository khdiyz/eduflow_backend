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
	Branch
}

func NewServices(repo *repository.Repository, storage *storage.Storage, cfg *config.Config, logger *logger.Logger) *Service {
	return &Service{
		Role:          NewRoleService(repo),
		Authorization: NewAuthService(repo, cfg),
		School:        NewSchoolService(repo),
		Branch:        NewBranchService(repo),
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
	GetSchool(id uuid.UUID) (models.School, error)
	Update(input models.UpdateSchool) error
	Delete(id uuid.UUID) error
}

type Branch interface {
	CreateBranch(input models.CreateBranch) (uuid.UUID, error)
	GetBranches(filter models.BranchFilter) ([]models.Branch, int, error)
	GetBranch(schoolId, branchId uuid.UUID) (models.Branch, error)
	UpdateBranch(input models.UpdateBranch) error
	DeleteBranch(schoolId, branchId uuid.UUID) error
}
