package domain

import (
	"fmt"
	"testing"
)

func TestNewEmployee_Valid(t *testing.T) {
	e, err := NewEmployee("Иван Иванов", "+77015556677", "Алматы")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if e.Fio != "Иван Иванов" { // вместо e.Fio
		t.Errorf("expected name to be Иван Иванов, got %s", e.Fio)
	}
	if e.City != "Алматы" {
		t.Errorf("expected city to be Алматы, got %s", e.City)
	}
}


func TestNewEmployee_InvalidPhone(t *testing.T) {
	_, err := NewEmployee("Иван Иванов", "12345", "Алматы")
	if err == nil {
		t.Fatal("expected error for invalid phone, got nil")
	}
}

func TestNewEmployee_InvalidName(t *testing.T) {
	_, err := NewEmployee("Иван", "+77015556677", "Алматы")
	fmt.Println("Test")
	if err == nil {
		t.Fatal("expected error for invalid full name, got nil")
	}
}
