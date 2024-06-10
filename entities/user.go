package entities

import (
	"time"
)

type User struct {
	ID           int
	EmployeeId   int
	FirstName    string
	LastName     string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}