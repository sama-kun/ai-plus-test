package domain

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

type Employee struct {
	Id        int        `json:"id" example:"1"`
	Fio      string     `json:"fio" example:"Иван Иванов"`
	Phone     string     `json:"phone" example:"+77071234567"`
	City      string     `json:"city" example:"Москва"`
	IsDeleted bool       `json:"isDeleted,omitempty" example:"false"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" example:"null"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty" example:"2025-09-16T12:00:00Z"`
	CreatedAt time.Time  `json:"createdAt,omitempty" example:"2025-09-10T09:30:00Z"`
}


func NewEmployee(fio, phone, city string) (*Employee, error) {
	fio = strings.TrimSpace(fio)
	phone = strings.TrimSpace(phone)
	city = strings.TrimSpace(city)

	if err := validateName(fio); err != nil {
		return nil, err
	}
	if err := validatePhone(phone); err != nil {
		return nil, err
	}

	return &Employee{
		Fio:  fio,
		Phone: phone,
		City:  city,
	}, nil
}

func validatePhone(phone string) error {
	re := regexp.MustCompile(`^(?:\+7\d{10}|8\d{10})$`)
	if !re.MatchString(phone) {
		return errors.New("invalid phone format, expected +7XXXXXXXXXX or 8XXXXXXXXXX")
	}
	return nil
}

func validateName(name string) error {
	parts := strings.Fields(name)
	if len(parts) < 2 {
		return errors.New("full name must contain at least two words")
	}
	for _, p := range parts {
		if len([]rune(p)) < 2 {
			return errors.New("each part of the name must be at least 2 characters")
		}
	}
	return nil
}
