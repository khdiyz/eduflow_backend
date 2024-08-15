package repository

import (
	"eduflow/internal/model"
	"eduflow/pkg/helper"
	"eduflow/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type RoleRepo struct {
	db  *sqlx.DB
	log logger.Logger
}

func NewRoleRepo(db *sqlx.DB, log logger.Logger) *RoleRepo {
	return &RoleRepo{
		db:  db,
		log: log,
	}
}

func (r *RoleRepo) GetList(pagination *model.Pagination) ([]model.Role, error) {
	var (
		roles []model.Role
		err   error
	)

	countQuery := "SELECT count(id) FROM roles WHERE deleted_at IS NULL;"
	err = helper.GetListCount(r.db, &r.log, pagination, countQuery, nil)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}

	query := `
	SELECT
		id,
		name,
		COALESCE(description, '') AS description,
		created_at,
		updated_at
	FROM roles
	WHERE
		deleted_at IS NULL
	LIMIT $1 OFFSET $2;`

	if err = r.db.Select(&roles, query, pagination.Limit, pagination.Offset); err != nil {
		r.log.Error(err)
		return nil, err
	}

	return roles, nil
}

func (r *RoleRepo) GetById(id int64) (model.Role, error) {
	var role model.Role

	query := `
	SELECT
		id,
		name,
		COALESCE(description, '') AS description,
		created_at,
		updated_at
	FROM roles
	WHERE
		deleted_at IS NULL
		AND id = $1;`

	if err := r.db.Get(&role, query, id); err != nil {
		return model.Role{}, err
	}

	return role, nil
}
