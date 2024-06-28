package types

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

type Timesheet struct {
	Id         int
	EmployeeId int
	StartTime  time.Time
	EndTime    time.Time
	Date  	   time.Time
	CreatedAt  time.Time
	ModifiedAt time.Time
}

type Paycycle struct {
	Id int
	EmployeeId int
	Timesheets []Timesheet
}

type LeaveRequest struct {
	Id         int
	EmployeeId int
	Type       string
	StartDate  time.Time
	EndDate	   time.Time
	TotalDays  int
	IsApproved bool
	CreatedAt  time.Time
	ModifiedAt time.Time
}