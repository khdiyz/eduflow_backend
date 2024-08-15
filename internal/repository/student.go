package repository

import (
	"eduflow/internal/model"
	"eduflow/pkg/helper"
	"eduflow/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type StudentRepo struct {
	db  *sqlx.DB
	log logger.Logger
}

func NewStudentRepo(db *sqlx.DB, log logger.Logger) *StudentRepo {
	return &StudentRepo{
		db:  db,
		log: log,
	}
}

func (r *StudentRepo) Create(input model.StudentCreateRequest) (int64, error) {
	var id int64

	query := `
	INSERT INTO students (
		full_name,
		phone_number_1,
		phone_number_2,
		address
	) VALUES ($1, $2, $3, $4) RETURNING id;`

	if err := r.db.QueryRow(query,
		input.FullName,
		input.Phone1,
		input.Phone2,
		input.Address,
	).Scan(&id); err != nil {
		r.log.Error(err)
		return 0, err
	}

	return id, nil
}

func (r *StudentRepo) GetList(pagination *model.Pagination) ([]model.Student, error) {
	var students []model.Student

	countQuery := "SELECT count(id) FROM students WHERE deleted_at IS NULL;"
	err := helper.GetListCount(r.db, &r.log, pagination, countQuery, nil)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}

	query := `
	SELECT
		id,
		full_name,
		phone_number_1,
		COALESCE(phone_number_2, ''),
		address,
		created_at,
		updated_at
	FROM students
	WHERE 
		deleted_at IS NULL
	LIMIT $1 OFFSET $2;`

	rows, err := r.db.Query(query, pagination.Limit, pagination.Offset)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}

	for rows.Next() {
		var student model.Student
		if err = rows.Scan(
			&student.Id,
			&student.FullName,
			&student.Phone1,
			&student.Phone2,
			&student.Address,
			&student.CreatedAt,
			&student.UpdatedAt,
		); err != nil {
			r.log.Error(err)
			return nil, err
		}

		students = append(students, student)
	}

	return students, nil
}
