package repository

import (
	"eduflow/internal/models"
	"eduflow/internal/repository/postgres"
	"eduflow/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Role
}

func NewRepository(db *sqlx.DB, logger *logger.Logger) *Repository {
	return &Repository{
		Role: postgres.NewRoleRepository(db, logger),
	}
}

type Role interface {
	Create(role models.CreateRole) error
}
