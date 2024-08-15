package service

import (
	"eduflow/internal/constants"
	"eduflow/internal/model"
	"eduflow/internal/repository"
	"eduflow/pkg/logger"
	"eduflow/pkg/response"
	"errors"

	"google.golang.org/grpc/codes"
)

type GroupService struct {
	repo repository.Repository
	log  logger.Logger
}

func NewGroupService(repo repository.Repository, log logger.Logger) *GroupService {
	return &GroupService{
		repo: repo,
		log:  log,
	}
}

func (s *GroupService) Create(input model.GroupCreateRequest) (int64, error) {
	_, err := s.repo.Course.GetById(input.CourseId)
	if err != nil {
		return 0, response.ServiceError(err, codes.Internal)
	}

	teacher, err := s.repo.User.GetById(input.TeacherId)
	if err != nil {
		return 0, response.ServiceError(err, codes.Internal)
	}

	if teacher.RoleName != constants.RoleMentor {
		return 0, response.ServiceError(errors.New("user must be teacher role"), codes.InvalidArgument)
	}

	id, err := s.repo.Group.Create(input)
	if err != nil {
		return 0, response.ServiceError(err, codes.Internal)
	}

	return id, nil
}

func (s *GroupService) GetList(pagination *model.Pagination) ([]model.Group, error) {
	groups, err := s.repo.Group.GetList(pagination)
	if err != nil {
		return nil, response.ServiceError(err, codes.Internal)
	}

	return groups, nil
}

func (s *GroupService) GetById(id int64) (model.Group, error) {
	group, err := s.repo.Group.GetById(id)
	if err != nil {
		return model.Group{}, response.ServiceError(err, codes.Internal)
	}

	return group, nil
}

func (s *GroupService) Update(input model.GroupUpdateRequest) error {
	_, err := s.repo.Course.GetById(input.CourseId)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}

	teacher, err := s.repo.User.GetById(input.TeacherId)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}

	if teacher.RoleName != constants.RoleMentor {
		return response.ServiceError(errors.New("user must be teacher role"), codes.InvalidArgument)
	}

	if err = s.repo.Group.Update(input); err != nil {
		return response.ServiceError(err, codes.Internal)
	}

	return nil
}

func (s *GroupService) Delete(id int64) error {
	err := s.repo.Group.Delete(id)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}

	return nil
}
