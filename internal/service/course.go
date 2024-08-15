package service

import (
	"eduflow/config"
	"eduflow/internal/model"
	"eduflow/internal/repository"
	"eduflow/pkg/helper"
	"eduflow/pkg/logger"
	"eduflow/pkg/response"

	"google.golang.org/grpc/codes"
)

type CourseService struct {
	repo repository.Repository
	log  logger.Logger
	cfg  config.Config
}

func NewCourseService(repo repository.Repository, log logger.Logger, cfg config.Config) *CourseService {
	return &CourseService{
		repo: repo,
		log:  log,
		cfg:  cfg,
	}
}

func (s *CourseService) Create(input model.CourseCreateRequest) (int64, error) {
	id, err := s.repo.Course.Create(input)
	if err != nil {
		return 0, response.ServiceError(err, codes.Internal)
	}

	return id, nil
}

func (s *CourseService) GetList(pagination *model.Pagination) ([]model.Course, error) {
	courses, err := s.repo.Course.GetList(pagination)
	if err != nil {
		return nil, response.ServiceError(err, codes.Internal)
	}

	for i := range courses {
		courses[i].Photo = helper.GenerateLink(s.cfg, courses[i].Photo)
	}

	return courses, nil
}

func (s *CourseService) GetById(id int64) (model.Course, error) {
	course, err := s.repo.Course.GetById(id)
	if err != nil {
		return model.Course{}, response.ServiceError(err, codes.Internal)
	}

	course.Photo = helper.GenerateLink(s.cfg, course.Photo)

	return course, nil
}

func (s *CourseService) Update(input model.CourseUpdateRequest) error {
	err := s.repo.Course.Update(input)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}

	return nil
}

func (s *CourseService) Delete(id int64) error {
	err := s.repo.Course.Delete(id)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}

	return nil
}
