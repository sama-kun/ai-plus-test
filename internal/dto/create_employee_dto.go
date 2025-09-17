package dto

type CreateEmployeeDTO struct {
	Name  string `json:"name" validate:"required"`
	Phone string `json:"phone" validate:"required"`
	City  string `json:"city" validate:"required"`
}
