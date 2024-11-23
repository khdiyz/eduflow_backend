package service

import (
	"eduflow/internal/models"
	"eduflow/internal/repository"
	"errors"
	"slices"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
)

type SchoolService struct {
	repo *repository.Repository
}

func NewSchoolService(repo *repository.Repository) *SchoolService {
	return &SchoolService{
		repo: repo,
	}
}

var currencyArr = []string{"UZS", "USD", "RUB"}

func (s *SchoolService) Create(input models.CreateSchool) (uuid.UUID, error) {
	input.Status = true

	if !slices.Contains(currencyArr, input.Currency) {
		return uuid.Nil, serviceError(errors.New("invalid currency value"), codes.InvalidArgument)
	}

	err := validateTimezone(input.Timezone)
	if err != nil {
		return uuid.Nil, serviceError(err, codes.InvalidArgument)
	}

	id, err := s.repo.School.Create(input)
	if err != nil {
		return uuid.Nil, serviceError(err, codes.Internal)
	}

	return id, nil
}

func validateTimezone(timezone string) error {
	_, err := time.LoadLocation(timezone)
	if err != nil {
		return errors.New("invalid timezone")
	}
	return nil
}

func (s *SchoolService) GetListSchool(filter models.SchoolFilter) ([]models.School, int, error) {
	schools, total, err := s.repo.School.GetList(filter)
	if err != nil {
		return nil, 0, serviceError(err, codes.Internal)
	}

	return schools, total, nil
}
