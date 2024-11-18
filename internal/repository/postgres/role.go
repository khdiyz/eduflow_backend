package postgres

import (
	"eduflow/internal/models"
	"eduflow/pkg/logger"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type RoleRepository struct {
	db     *sqlx.DB
	logger *logger.Logger
}

func NewRoleRepository(db *sqlx.DB, logger *logger.Logger) *RoleRepository {
	return &RoleRepository{
		db:     db,
		logger: logger,
	}
}

func (r *RoleRepository) Create(role models.CreateRole) error {
	id := uuid.New()

	query := `
	INSERT INTO roles (
		id,
		name,
		description
	) VALUES ($1, $2, $3);`

	name, err := json.Marshal(role.Name)
	if err != nil {
		r.logger.Error(err)
		return err
	}

	description, err := json.Marshal(role.Description)
	if err != nil {
		r.logger.Error(err)
		return err
	}

	_, err = r.db.Exec(query, id, name, description)
	if err != nil {
		r.logger.Error(err)
		return err
	}

	return nil
}
