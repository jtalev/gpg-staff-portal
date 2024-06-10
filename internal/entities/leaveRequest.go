package entities

import "time"

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