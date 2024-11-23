package repository

import (
	"eduflow/internal/models"
	"eduflow/internal/repository/postgres"
	"eduflow/pkg/logger"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Role
	User
	School
}

func NewRepository(db *sqlx.DB, logger *logger.Logger) *Repository {
	return &Repository{
		Role:   postgres.NewRoleRepository(db, logger),
		User:   postgres.NewUserRepository(db, logger),
		School: postgres.NewSchoolRepository(db, logger),
	}
}

type Role interface {
	Create(role models.CreateRole) error
}

type User interface {
	Create(input models.CreateUser) (uuid.UUID, error)
	GetList(filter models.UserFilter) ([]models.User, int, error)
	GetById(id uuid.UUID) (models.User, error)
	GetByUsername(username string) (models.User, error)
}

type School interface {
	Create(input models.CreateSchool) (uuid.UUID, error)
	GetList(filter models.SchoolFilter) ([]models.School, int, error)
}
