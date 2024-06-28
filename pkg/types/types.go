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

type Login struct {
	Email    string
	Password string
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

type ValidateIntTest struct {
	Expected Result
	Value    int
	Field    string
}

type ValidateStringTest struct {
	Expected Result
	Value    string
	Field    string
}

type ValidateTimeTest struct {
	Expected Result
	Start    time.Time
	End    time.Time
	Date    time.Time
	Field    string
}

type Result struct {
	IsValid bool
	Error   string
}
