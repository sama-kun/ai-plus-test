package handler

import (
	"context"
	"net/http"

	"github.com/sama-kun/ai-plus-test/internal/dto"
	"github.com/sama-kun/ai-plus-test/internal/lib/middleware"
	"github.com/sama-kun/ai-plus-test/internal/service"
)

type EmployeeHandler struct {
	svc *service.EmployeeService
}

func NewEmployeeHandler(s *service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{svc: s}
}

// @Summary Create employee
// @Description Adds a new employee
// @Accept  json
// @Produce  json
// @Param employee body Employee true "Employee info"
// @Success 201 {object} Employee
// @Failure 400 {object} map[string]string
// @Router /employee [post]
func (h *EmployeeHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateEmployeeDTO

	if err := middleware.DecodeJSON(r, &req); err != nil {
		middleware.ErrorHandler(w, http.StatusBadRequest, err, "Invalid request body")
		return
	}
	resp, err := h.svc.Create(context.Background(), req.Name, req.Phone, req.City)
	if err != nil {
		middleware.ErrorHandler(w, http.StatusInternalServerError, err, "Registration employee failed")
		return
	}
	middleware.JSONResponse(w, http.StatusCreated, resp)
}
