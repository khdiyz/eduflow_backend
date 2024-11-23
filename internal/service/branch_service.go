package service

import (
	"eduflow/internal/models"
	"eduflow/internal/repository"
	"errors"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
)

type BranchService struct {
	repo *repository.Repository
}

func NewBranchService(repo *repository.Repository) *BranchService {
	return &BranchService{
		repo: repo,
	}
}

func (s *BranchService) CreateBranch(input models.CreateBranch) (uuid.UUID, error) {
	input.Status = true

	id, err := s.repo.Branch.Create(input)
	if err != nil {
		return uuid.Nil, serviceError(err, codes.Internal)
	}

	return id, nil
}

func (s *BranchService) GetBranches(filter models.BranchFilter) ([]models.Branch, int, error) {
	branches, total, err := s.repo.Branch.GetList(filter)
	if err != nil {
		return nil, 0, serviceError(err, codes.Internal)
	}

	return branches, total, nil
}

func (s *BranchService) GetBranch(schoolId, branchId uuid.UUID) (models.Branch, error) {
	branch, err := s.repo.Branch.GetById(branchId)
	if err != nil {
		return models.Branch{}, serviceError(err, codes.Internal)
	}

	if branch.SchoolId != schoolId {
		return models.Branch{}, serviceError(errors.New("invalid school id"), codes.InvalidArgument)
	}

	return branch, nil
}

func (s *BranchService) UpdateBranch(input models.UpdateBranch) error {
	branch, err := s.repo.Branch.GetById(input.Id)
	if err != nil {
		return serviceError(err, codes.Internal)
	}

	if input.SchoolId != branch.SchoolId {
		return serviceError(errors.New("invalid school id"), codes.InvalidArgument)
	}

	err = s.repo.Branch.Update(input)
	if err != nil {
		return serviceError(err, codes.Internal)
	}

	return nil
}

func (s *BranchService) DeleteBranch(schoolId, branchId uuid.UUID) error {
	branch, err := s.repo.Branch.GetById(branchId)
	if err != nil {
		return serviceError(err, codes.Internal)
	}

	if schoolId != branch.SchoolId {
		return serviceError(errors.New("invalid school id"), codes.InvalidArgument)
	}

	err = s.repo.Branch.Delete(branchId)
	if err != nil {
		return serviceError(err, codes.Internal)
	}

	return nil
}
