package repository

import (
	"eduflow/internal/constants"
	"eduflow/internal/model"
	"eduflow/pkg/helper"
	"eduflow/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type GroupRepo struct {
	db  *sqlx.DB
	log logger.Logger
}

func NewGroupRepo(db *sqlx.DB, log logger.Logger) *GroupRepo {
	return &GroupRepo{
		db:  db,
		log: log,
	}
}

func (r *GroupRepo) Create(input model.GroupCreateRequest) (int64, error) {
	var id int64

	query := `
	INSERT INTO groups (
		name,
		start_date,
		course_id,
		teacher_id
	) VALUES ($1, $2, $3, $4) RETURNING id;`

	err := r.db.QueryRow(query, input.Name, input.StartDate, input.CourseId, input.TeacherId).Scan(&id)
	if err != nil {
		r.log.Error(err)
		return 0, err
	}

	return id, nil
}

func (r *GroupRepo) GetList(pagination *model.Pagination) ([]model.Group, error) {
	var groups []model.Group

	countQuery := "SELECT count(id) FROM groups WHERE deleted_at IS NULL;"
	err := helper.GetListCount(r.db, &r.log, pagination, countQuery, nil)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}

	query := `
	SELECT
		g.id,
		g.name,
		g.start_date,
		g.end_date,
		c.id,
		c.name,
		u.id,
		u.full_name,
		g.created_at,
		g.updated_at
	FROM groups g 
	JOIN courses c ON g.course_id = c.id
	JOIN users u ON g.teacher_id = u.id
	WHERE
		g.deleted_at IS NULL
		AND c.deleted_at IS NULL
		AND u.deleted_at IS NULL
	LIMIT $1 OFFSET $2;`

	rows, err := r.db.Query(query, pagination.Limit, pagination.Offset)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}

	for rows.Next() {
		var group model.Group
		if err = rows.Scan(
			&group.Id,
			&group.Name,
			&group.StartDate,
			&group.EndDate,
			&group.Course.Id,
			&group.Course.Name,
			&group.Teacher.Id,
			&group.Teacher.FullName,
			&group.CreatedAt,
			&group.UpdatedAt,
		); err != nil {
			r.log.Error(err)
			return nil, err
		}

		groups = append(groups, group)
	}

	return groups, nil
}

func (r *GroupRepo) GetById(id int64) (model.Group, error) {
	var group model.Group

	query := `
	SELECT
		g.id,
		g.name,
		g.start_date,
		g.end_date,
		c.id,
		c.name,
		u.id,
		u.full_name,
		g.created_at,
		g.updated_at
	FROM groups g 
	JOIN courses c ON g.course_id = c.id
	JOIN users u ON g.teacher_id = u.id
	WHERE
		g.deleted_at IS NULL
		AND c.deleted_at IS NULL
		AND u.deleted_at IS NULL
		AND g.id = $1;`

	if err := r.db.QueryRow(query, id).Scan(
		&group.Id,
		&group.Name,
		&group.StartDate,
		&group.EndDate,
		&group.Course.Id,
		&group.Course.Name,
		&group.Teacher.Id,
		&group.Teacher.FullName,
		&group.CreatedAt,
		&group.UpdatedAt,
	); err != nil {
		return model.Group{}, err
	}

	return group, nil
}

func (r *GroupRepo) Update(input model.GroupUpdateRequest) error {
	query := `
	UPDATE groups
	SET
		name = $2,
		start_date = $3,
		end_date = $4,
		course_id = $5,
		teacher_id = $6,
		updated_at = now()
	WHERE 
		id = $1
		AND deleted_at IS NULL;`

	result, err := r.db.Exec(query,
		input.Id,
		input.Name,
		input.StartDate,
		input.EndDate,
		input.CourseId,
		input.TeacherId,
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

func (r *GroupRepo) Delete(id int64) error {
	query := `
	UPDATE groups
	SET
		deleted_at = now()
	WHERE 
		id = $1
		AND deleted_at IS NULL;`

	row, err := r.db.Exec(query, id)
	if err != nil {
		r.log.Error(err)
		return err
	}
	rowAffected, err := row.RowsAffected()
	if err != nil {
		r.log.Error(err)
		return err
	}
	if rowAffected == 0 {
		return constants.ErrNoRowsAffected
	}

	return nil
}
