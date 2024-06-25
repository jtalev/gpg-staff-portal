package entities

import (
	"time"
)

type User struct {
	Uid          int
	EmployeeId   int
	FirstName    string
	LastName     string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	ModifiedAt   time.Time
}