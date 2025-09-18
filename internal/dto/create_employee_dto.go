package dto

type CreateEmployeeDTO struct {
	Fio  string `json:"fio" validate:"required" example:"Sama Seriknur"`
	Phone string `json:"phone" validate:"required" example:"+77071231212"`
	City  string `json:"city" validate:"required" example:"Almaty"`
}
