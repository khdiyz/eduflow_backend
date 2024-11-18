package service

import (
	"eduflow/internal/models"
	"eduflow/internal/repository"

	"google.golang.org/grpc/codes"
)

type RoleService struct {
	repo *repository.Repository
}

func NewRoleService(repo *repository.Repository) *RoleService {
	return &RoleService{
		repo: repo,
	}
}

func (s *RoleService) Create(role models.CreateRole) error {
	if err := s.repo.Create(role); err != nil {
		return serviceError(err, codes.Internal)
	}

	return nil
}
