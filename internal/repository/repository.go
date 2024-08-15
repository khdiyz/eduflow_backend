package repository

import (
	"eduflow/internal/model"
	"eduflow/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	User
	Role
	Course
	Group
<<<<<<< HEAD
	Student
=======
>>>>>>> bd4e226 (initial)
}

func NewRepository(db *sqlx.DB, log logger.Logger) *Repository {
	return &Repository{
<<<<<<< HEAD
		User:    NewUserRepo(db, log),
		Role:    NewRoleRepo(db, log),
		Course:  NewCourseRepo(db, log),
		Group:   NewGroupRepo(db, log),
		Student: NewStudentRepo(db, log),
=======
		User:   NewUserRepo(db, log),
		Role:   NewRoleRepo(db, log),
		Course: NewCourseRepo(db, log),
		Group:  NewGroupRepo(db, log),
>>>>>>> bd4e226 (initial)
	}
}

type User interface {
	Create(input model.UserCreateRequest) (int64, error)
	GetByUsername(username string) (model.User, error)
	GetById(id int64) (model.User, error)
	GetList(pagination *model.Pagination, filters map[string]interface{}) ([]model.User, error)
	Update(user model.UserUpdateRequest) error
	DeleteById(id int64) error
}

type Role interface {
	GetList(pagination *model.Pagination) ([]model.Role, error)
	GetById(id int64) (model.Role, error)
}

type Course interface {
	Create(input model.CourseCreateRequest) (int64, error)
	GetList(pagination *model.Pagination) ([]model.Course, error)
	GetById(id int64) (model.Course, error)
	Update(input model.CourseUpdateRequest) error
	Delete(id int64) error
}

type Group interface {
	Create(input model.GroupCreateRequest) (int64, error)
	GetList(pagination *model.Pagination) ([]model.Group, error)
	GetById(id int64) (model.Group, error)
	Update(input model.GroupUpdateRequest) error
	Delete(id int64) error
}
<<<<<<< HEAD

type Student interface {
	Create(input model.StudentCreateRequest) (int64, error)
	GetList(pagination *model.Pagination) ([]model.Student, error)
}
=======
>>>>>>> bd4e226 (initial)
