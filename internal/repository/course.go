package repository

import (
	"eduflow/internal/constants"
	"eduflow/internal/model"
	"eduflow/pkg/helper"
	"eduflow/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type CourseRepo struct {
	db  *sqlx.DB
	log logger.Logger
}

func NewCourseRepo(db *sqlx.DB, log logger.Logger) *CourseRepo {
	return &CourseRepo{
		db:  db,
		log: log,
	}
}

func (r *CourseRepo) Create(input model.CourseCreateRequest) (int64, error) {
	var id int64

	query := `
	INSERT INTO courses (
		name,
		description,
		photo
	) VALUES ($1, $2, $3) RETURNING id;`

	err := r.db.QueryRow(query, input.Name, input.Description, input.Photo).Scan(&id)
	if err != nil {
		r.log.Error(err)
		return 0, err
	}

	return id, nil
}

func (r *CourseRepo) GetList(pagination *model.Pagination) ([]model.Course, error) {
	var courses []model.Course

	countQuery := "SELECT count(id) FROM courses WHERE deleted_at IS NULL;"
	err := helper.GetListCount(r.db, &r.log, pagination, countQuery, nil)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}

	query := `
	SELECT
		id,
		name,
		COALESCE(description, '') AS description,
		COALESCE(photo, '') AS photo,
		created_at,
		updated_at
	FROM courses
	WHERE 
		deleted_at IS NULL
	LIMIT $1 OFFSET $2;`

	rows, err := r.db.Query(query, pagination.Limit, pagination.Offset)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}

	for rows.Next() {
		var course model.Course
		if err = rows.Scan(
			&course.Id,
			&course.Name,
			&course.Description,
			&course.Photo,
			&course.CreatedAt,
			&course.UpdatedAt,
		); err != nil {
			r.log.Error(err)
			return nil, err
		}

		courses = append(courses, course)
	}

	return courses, nil
}

func (r *CourseRepo) GetById(id int64) (model.Course, error) {
	var course model.Course

	query := `
	SELECT
		id,
		name,
		COALESCE(description, '') AS description,
		COALESCE(photo, '') AS photo,
		created_at,
		updated_at
	FROM courses
	WHERE
		deleted_at IS NULL
		AND id = $1;`

	if err := r.db.QueryRow(query, id).Scan(
		&course.Id,
		&course.Name,
		&course.Description,
		&course.Photo,
		&course.CreatedAt,
		&course.UpdatedAt,
	); err != nil {
		return model.Course{}, err
	}

	return course, nil
}

func (r *CourseRepo) Update(input model.CourseUpdateRequest) error {
	query := `
	UPDATE courses
	SET
		name = $2,
		description = $3,
		photo = $4,
		updated_at = now()
	WHERE 
		id = $1
		AND deleted_at IS NULL;`

	result, err := r.db.Exec(query, input.Id, input.Name, input.Description, input.Photo)
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

func (r *CourseRepo) Delete(id int64) error {
	query := `
	UPDATE courses
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
