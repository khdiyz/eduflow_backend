package service

import (
	"database/sql"
	"eduflow/config"
	"eduflow/internal/model"
	"eduflow/internal/repository"
	"eduflow/pkg/helper"
	"eduflow/pkg/logger"
	"eduflow/pkg/response"
	"errors"

	"google.golang.org/grpc/codes"
)

type UserService struct {
	repo repository.Repository
	log  logger.Logger
	cfg  config.Config
}

func NewUserService(repo repository.Repository, log logger.Logger, cfg config.Config) *UserService {
	return &UserService{
		repo: repo,
		log:  log,
		cfg:  cfg,
	}
}

func (s *UserService) Create(request model.UserCreateRequest) (int64, error) {
	_, err := s.repo.Role.GetById(request.RoleId)
	if err != nil {
		return 0, response.ServiceError(errors.New("role with this id does not exist"), codes.InvalidArgument)
	}

	_, err = s.repo.User.GetByUsername(request.Username)
	if err != nil && err != sql.ErrNoRows {
		return 0, response.ServiceError(err, codes.Internal)
	} else if err == nil {
		return 0, response.ServiceError(errors.New("user exists with this username"), codes.InvalidArgument)
	}

	hash, err := helper.GenerateHash(request.Password)
	if err != nil {
		return 0, response.ServiceError(err, codes.Internal)
	}
	request.Password = hash

	id, err := s.repo.User.Create(request)
	if err != nil {
		return 0, response.ServiceError(err, codes.Internal)
	}

	return id, nil
}

func (s *UserService) GetById(id int64) (model.User, error) {
	user, err := s.repo.User.GetById(id)
	if err != nil {
		return model.User{}, response.ServiceError(err, codes.Internal)
	}

	user.Photo = helper.GenerateLink(s.cfg, user.Photo)

	return user, nil
}

func (s *UserService) GetList(pagination *model.Pagination, filters map[string]interface{}) ([]model.User, error) {
	users, err := s.repo.User.GetList(pagination, filters)
	if err != nil {
		return nil, response.ServiceError(err, codes.Internal)
	}

	for i := range users {
		users[i].Photo = helper.GenerateLink(s.cfg, users[i].Photo)
	}

	return users, nil
}

func (s *UserService) Update(user model.UserUpdateRequest) error {
	err := s.repo.User.Update(user)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}

	return nil
}

func (s *UserService) DeleteById(id int64) error {
	err := s.repo.User.DeleteById(id)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}

	return nil
}
