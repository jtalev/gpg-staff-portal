package entities

import (
	"time"
)

type LeaveRequest struct {
	ID         	int
	EmployeeId 	int
	StartDate   time.Time
	EndDate		time.Time
	LeaveType	string
	HrsPayable	float32
	IsAccepted 	bool
}