package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sama-kun/ai-plus-test/internal/domain"
	"github.com/sama-kun/ai-plus-test/internal/handler"
	"github.com/sama-kun/ai-plus-test/internal/service"
)

type mockRepo struct {
	employees []domain.Employee
}

func (m *mockRepo) Save(ctx context.Context, e *domain.Employee) (int, error) {
	e.Id = len(m.employees) + 1
	m.employees = append(m.employees, *e)
	return e.Id, nil
}

func (m *mockRepo) FindAll(ctx context.Context) ([]*domain.Employee, error) {
	var result []*domain.Employee
	for i := range m.employees {
		result = append(result, &m.employees[i])
	}
	return result, nil
}

func TestCreateEmployeeHandler(t *testing.T) {
	repo := &mockRepo{}
	svc := service.NewEmployeeService(repo)
	h := handler.NewEmployeeHandler(svc)

	employee := map[string]string{
		"fio":  "Иван Иванов",
		"phone": "+77015556677",
		"city":  "Алматы",
	}
	body, _ := json.Marshal(employee)

	req := httptest.NewRequest(http.MethodPost, "/employee", bytes.NewReader(body))
	w := httptest.NewRecorder()

	h.CreateEmployee(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", w.Code)
	}
}
