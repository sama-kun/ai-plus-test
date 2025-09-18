package service

import (
	"context"

	"github.com/sama-kun/ai-plus-test/internal/domain"
	"github.com/sama-kun/ai-plus-test/internal/dto"
	"github.com/sama-kun/ai-plus-test/internal/repository"
)

type EmployeeService struct {
	repo repository.PostgresEmployeeRepoInterface
}

func NewEmployeeService(r repository.PostgresEmployeeRepoInterface) *EmployeeService {
	return &EmployeeService{repo: r}
}

func (s *EmployeeService) Create(ctx context.Context, fio, phone, city string) (*dto.CreateEmployeeResponse, error) {
	e, err := domain.NewEmployee(fio, phone, city)
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

func (s *EmployeeService) FindAll(ctx context.Context) ([]*domain.Employee, error) {
	employees, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var result []*domain.Employee
	for _, e := range employees {
		result = append(result, &domain.Employee{
			Id:        e.Id,
			Fio:      e.Fio,
			Phone:     e.Phone,
			City:      e.City,
			IsDeleted: e.IsDeleted,
			DeletedAt: e.DeletedAt,
			UpdatedAt: e.UpdatedAt,
			CreatedAt: e.CreatedAt,
		})
	}

	return result, nil
}