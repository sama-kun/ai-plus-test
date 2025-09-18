package repository

import (
	"context"
	"fmt"

	"github.com/sama-kun/ai-plus-test/internal/domain"
	"github.com/sama-kun/ai-plus-test/internal/storage"
)

type PostgresEmployeeRepoInterface interface {
	Save(ctx context.Context, e *domain.Employee) (int, error)
	FindAll(ctx context.Context) ([]*domain.Employee, error)
}

type PostgresEmployeeRepo struct {
	db *storage.Postgres
}

func NewPostgresEmployeeRepo(db *storage.Postgres) PostgresEmployeeRepoInterface {
	return &PostgresEmployeeRepo{db: db}
}

func (r *PostgresEmployeeRepo) Save(ctx context.Context, e *domain.Employee) (int, error) {
	const fn = "repository.employee.Save"

	const query = `
	INSERT INTO employees (name, phone, city) 
	VALUES ($1, $2, $3)
	RETURNING id`
	var id int
	err := r.db.Conn().QueryRow(
		ctx,
		query,
		e.Fio, e.Phone, e.City,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", fn, err)
	}
	return id, err
}

func (r *PostgresEmployeeRepo) FindAll(ctx context.Context) ([]*domain.Employee, error) {
	const fn = "repository.employee.FindAll"

	const query = `
		SELECT id, name, phone, city, is_deleted, deleted_at, updated_at, created_at
		FROM employees
		WHERE is_deleted = false
		ORDER BY created_at DESC`

	rows, err := r.db.Conn().Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	defer rows.Close()

	var employees []*domain.Employee
	for rows.Next() {
		var e domain.Employee
		err := rows.Scan(
			&e.Id,
			&e.Fio,
			&e.Phone,
			&e.City,
			&e.IsDeleted,
			&e.DeletedAt,
			&e.UpdatedAt,
			&e.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", fn, err)
		}
		employees = append(employees, &e)
	}

	return employees, nil
}