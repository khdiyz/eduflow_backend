package repository

import (
	"eduflow/internal/constants"
	"eduflow/internal/model"
	"eduflow/pkg/helper"
	"eduflow/pkg/logger"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db  *sqlx.DB
	log logger.Logger
}

func NewUserRepo(db *sqlx.DB, log logger.Logger) *UserRepo {
	return &UserRepo{
		db:  db,
		log: log,
	}
}

func (r *UserRepo) Create(input model.UserCreateRequest) (int64, error) {
	var id int64

	query := `
	INSERT INTO users (
		full_name,
		phone_number,
		birth_date,
		photo,
		role_id,
		username,
		password
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7 
	) RETURNING id;`

	if err := r.db.QueryRow(query,
		input.FullName,
		input.PhoneNumber,
		input.BirthDate,
		input.Photo,
		input.RoleId,
		input.Username,
		input.Password,
	).Scan(&id); err != nil {
		r.log.Error(err)
		return 0, err
	}

	return id, nil
}

func (r *UserRepo) GetByUsername(username string) (model.User, error) {
	var user model.User

	query := `
	SELECT
		u.id,
		u.full_name,
		u.phone_number,
		u.birth_date,
		u.photo,
		u.role_id,
		r.name,
		u.username,
		u.password,
		u.created_at,
		u.updated_at
	FROM users u 
	JOIN roles r ON u.role_id = r.id
	WHERE 
		u.username = $1
		AND u.deleted_at IS NULL
		AND r.deleted_at IS NULL;`

	if err := r.db.QueryRow(query, username).Scan(
		&user.Id,
		&user.FullName,
		&user.PhoneNumber,
		&user.BirthDate,
		&user.Photo,
		&user.RoleId,
		&user.RoleName,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		r.log.Error(err)
		return model.User{}, err
	}

	return user, nil
}

func (r *UserRepo) GetById(id int64) (model.User, error) {
	var user model.User

	query := `
	SELECT
		u.id,
		u.full_name,
		u.phone_number,
		u.birth_date,
		COALESCE(u.photo, ''),
		u.role_id,
		r.name,
		u.username,
		u.password,
		u.created_at,
		u.updated_at
	FROM users u 
	JOIN roles r ON u.role_id = r.id
	WHERE 
		u.id = $1
		AND u.deleted_at IS NULL
		AND r.deleted_at IS NULL;`

	if err := r.db.QueryRow(query, id).Scan(
		&user.Id,
		&user.FullName,
		&user.PhoneNumber,
		&user.BirthDate,
		&user.Photo,
		&user.RoleId,
		&user.RoleName,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		r.log.Error(err)
		return model.User{}, err
	}

	return user, nil
}

func (r *UserRepo) GetList(pagination *model.Pagination, filters map[string]interface{}) ([]model.User, error) {
	var (
		users []model.User
		err   error
	)

	countQuery := "SELECT count(id) FROM users WHERE deleted_at IS NULL"

	query := `
	SELECT
		u.id,
		u.full_name,
		u.phone_number,
		u.birth_date,
		COALESCE(u.photo, ''),
		u.role_id,
		r.name,
		u.username,
		u.password,
		u.created_at,
		u.updated_at
	FROM users u 
	JOIN roles r ON u.role_id = r.id
	WHERE
		u.deleted_at IS NULL
		AND r.deleted_at IS NULL `

	var filterClauses []string
	var args []interface{}
	var counter int

	if roleId, ok := filters["role_id"]; ok {
		counter++
		filterClauses = append(filterClauses, "role_id = $"+strconv.Itoa(counter))
		args = append(args, roleId)
	}

	if len(filterClauses) > 0 {
		countQuery += " AND " + strings.Join(filterClauses, " AND ")
		query += " AND " + strings.Join(filterClauses, " AND ")
	}

	query += " LIMIT $" + strconv.Itoa(counter+1) + " OFFSET $" + strconv.Itoa(counter+2)

	args = append(args, pagination.Limit, pagination.Offset)

	err = helper.GetListCount(r.db, &r.log, pagination, countQuery, args[:len(args)-2])
	if err != nil {
		r.log.Error(err)
		return nil, err
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}

	for rows.Next() {
		var user model.User
		if err = rows.Scan(
			&user.Id,
			&user.FullName,
			&user.PhoneNumber,
			&user.BirthDate,
			&user.Photo,
			&user.RoleId,
			&user.RoleName,
			&user.Username,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			r.log.Error(err)
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepo) Update(user model.UserUpdateRequest) error {
	query := `
	UPDATE users
	SET 
		full_name = $2,
		phone_number = $3,
		birth_date = $4,
		photo = $5,
		role_id = $6,
		username = $7
	WHERE
		id = $1
		AND deleted_at IS NULL;`

	result, err := r.db.Exec(query,
		user.Id,
		user.FullName,
		user.PhoneNumber,
		user.BirthDate,
		user.Photo,
		user.RoleId,
		user.Username,
	)
	if err != nil {
		r.log.Error(err)
		return err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		r.log.Error(err)
		return err
	}
	if rowAffected == 0 {
		return constants.ErrNoRowsAffected
	}

	return nil
}

func (r *UserRepo) DeleteById(id int64) error {
	query := `
	UPDATE users
	SET
		deleted_at = now()
	WHERE
		id = $1
		AND deleted_at IS NULL;`

	result, err := r.db.Exec(query, id)
	if err != nil {
		r.log.Error(err)
		return err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		r.log.Error(err)
		return err
	}
	if rowAffected == 0 {
		return constants.ErrNoRowsAffected
	}

	return nil
}
