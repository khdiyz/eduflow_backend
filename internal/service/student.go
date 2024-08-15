package service

import (
	"eduflow/internal/model"
	"eduflow/internal/repository"
	"eduflow/pkg/logger"
	"eduflow/pkg/response"

	"google.golang.org/grpc/codes"
)

type StudentService struct {
	repo repository.Repository
	log  logger.Logger
}

func NewStudentService(repo repository.Repository, log logger.Logger) *StudentService {
	return &StudentService{
		repo: repo,
		log:  log,
	}
}

func (s *StudentService) Create(input model.StudentCreateRequest) (int64, error) {
	id, err := s.repo.Student.Create(input)
	if err != nil {
		return 0, response.ServiceError(err, codes.Internal)
	}

	return id, nil
}

func (s *StudentService) GetList(pagination *model.Pagination) ([]model.Student, error) {
	students, err := s.repo.Student.GetList(pagination)
	if err != nil {
		return nil, response.ServiceError(err, codes.Internal)
	}

	return students, nil
}
