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
	Email    	string
	Password 	string
}

type Timesheet struct {
	Id         	int
	EmployeeId 	int
	StartTime  	time.Time
	EndTime    	time.Time
	Date  	   	time.Time
	CreatedAt  	time.Time
	ModifiedAt 	time.Time
}

type Paycycle struct {
	Id 			int
	EmployeeId 	int
	StartDate	time.Time
	EndDate		time.Time
}

type LeaveRequest struct {
	Id         	int
	EmployeeId 	int
	Type       	string
	StartDate  	time.Time
	EndDate	   	time.Time
	TotalDays  	int
	IsApproved 	bool
	CreatedAt  	time.Time
	ModifiedAt 	time.Time
}

type ValidateIntTest struct {
	Expected 	ValidationResult
	Value    	int
	Field    	string
}

type ValidateStringTest struct {
	Expected 	ValidationResult
	Value    	string
	Field    	string
}

type ValidateTimeTest struct {
	Expected 	ValidationResult
	Start    	time.Time
	End    	 	time.Time
	Date     	time.Time
	Field    	string
}

type ValidateBoolTest struct {
	Expected 	ValidationResult
	Value 	 	bool
	Field	 	string
}

type ValidationResult struct {
	IsValid 	bool
	Error   	string
}
