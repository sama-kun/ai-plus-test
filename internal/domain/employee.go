package domain

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

type Employee struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Phone     string     `json:"phone"`
	City      string     `json:"city"`
	IsDeleted bool       `json:"isDeleted,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
}

func NewEmployee(name, phone, city string) (*Employee, error) {
	name = strings.TrimSpace(name)
	phone = strings.TrimSpace(phone)
	city = strings.TrimSpace(city)

	if err := validateName(name); err != nil {
		return nil, err
	}
	if err := validatePhone(phone); err != nil {
		return nil, err
	}

	return &Employee{
		Name:  name,
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
