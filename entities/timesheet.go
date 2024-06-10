package entities

import (
	"time"
)

type Timesheet struct {
	ID   int
	EmployeeId int
	Date time.Time
	StartTime time.Time
	EndTime time.Time
	CreatedAt time.Time
	ModifiedAt time.Time
}