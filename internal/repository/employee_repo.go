package repository

import (
	"context"
	"fmt"

	"github.com/sama-kun/ai-plus-test/internal/domain"
	"github.com/sama-kun/ai-plus-test/internal/storage"
)

type EmployeeRepository interface {
	Save(ctx context.Context, e *domain.Employee) (string, error)
}

type PostgresEmployeeRepo struct {
	db *storage.Postgres
}

func NewPostgresEmployeeRepo(db *storage.Postgres) EmployeeRepository {
	return &PostgresEmployeeRepo{db: db}
}

func (r *PostgresEmployeeRepo) Save(ctx context.Context, e *domain.Employee) (string, error) {
	const fn = "repository.employee.Save"

	const query = `
	INSERT INTO employees (name, phone, city) 
	VALUES ($1, $2, $3)
	RETURNING id`
	var id string
	err := r.db.Conn().QueryRow(
		ctx,
		query,
		e.Name, e.Phone, e.City,
	).Scan(id)
	if err != nil {
		return "", fmt.Errorf("%s: %w", fn, err)
	}
	return id, err
}
