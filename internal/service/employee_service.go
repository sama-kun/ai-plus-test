package service

import (
	"context"

	"github.com/sama-kun/ai-plus-test/internal/domain"
	"github.com/sama-kun/ai-plus-test/internal/dto"
	"github.com/sama-kun/ai-plus-test/internal/repository"
)

type EmployeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(r repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{repo: r}
}

func (s *EmployeeService) Create(ctx context.Context, name, phone, city string) (*dto.CreateEmployeeResponse, error) {
	e, err := domain.NewEmployee(name, phone, city)
	if err != nil {
		return nil, err
	}
	id, err := s.repo.Save(ctx, e)

	if err != nil {
		return nil, err
	}

	return &dto.CreateEmployeeResponse{
		Id: id,
	}, nil
}
