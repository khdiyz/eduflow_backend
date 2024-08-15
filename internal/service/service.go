package service

import (
	"eduflow/config"
	"eduflow/internal/model"
	"eduflow/internal/repository"
	"eduflow/internal/storage"
	"eduflow/pkg/logger"
	"io"
	"time"
)

type Service struct {
	Authorization
	User
	Minio
	Role
	Course
	Group
<<<<<<< HEAD
	Student
=======
>>>>>>> bd4e226 (initial)
}

func NewService(repos repository.Repository, strg storage.Storage, log logger.Logger, cfg config.Config) *Service {
	return &Service{
		Authorization: NewAuthService(repos, log),
		User:          NewUserService(repos, log, cfg),
		Minio:         NewMinioService(strg, log),
		Role:          NewRoleService(repos, log),
		Course:        NewCourseService(repos, log, cfg),
		Group:         NewGroupService(repos, log),
<<<<<<< HEAD
		Student:       NewStudentService(repos, log),
=======
>>>>>>> bd4e226 (initial)
	}
}

type Authorization interface {
	CreateToken(user model.User, tokenType string, expiresAt time.Time) (*model.Token, error)
	GenerateTokens(user model.User) (*model.Token, *model.Token, error)
	ParseToken(token string) (*jwtCustomClaim, error)
	Login(input model.LoginRequest) (*model.Token, *model.Token, error)
}

type User interface {
	Create(request model.UserCreateRequest) (int64, error)
	GetById(id int64) (model.User, error)
	GetList(pagination *model.Pagination, filters map[string]interface{}) ([]model.User, error)
	Update(user model.UserUpdateRequest) error
	DeleteById(id int64) error
}

type Minio interface {
	UploadImage(image io.Reader, imageSize int64, contextType string) (storage.File, error)
	UploadDoc(doc io.Reader, docSize int64, contextType string) (storage.File, error)
	UploadExcel(doc io.Reader, docSize int64, contextType string) (storage.File, error)
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
