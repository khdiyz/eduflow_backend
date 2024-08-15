package service

import (
	"eduflow/internal/model"
	"eduflow/internal/repository"
	"eduflow/pkg/logger"
	"eduflow/pkg/response"

	"google.golang.org/grpc/codes"
)

type RoleService struct {
	repo repository.Repository
	log  logger.Logger
}

func NewRoleService(repo repository.Repository, log logger.Logger) *RoleService {
	return &RoleService{
		repo: repo,
		log:  log,
	}
}

func (s *RoleService) GetList(pagination *model.Pagination) ([]model.Role, error) {
	roles, err := s.repo.Role.GetList(pagination)
	if err != nil {
		return nil, response.ServiceError(err, codes.Internal)
	}

	return roles, nil
}

func (s *RoleService) GetById(id int64) (model.Role, error) {
	role, err := s.repo.Role.GetById(id)
	if err != nil {
		return model.Role{}, response.ServiceError(err, codes.Internal)
	}

	return role, nil
}
