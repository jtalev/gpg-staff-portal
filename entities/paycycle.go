package entities

import (
	"time"
)

type PayCycle struct {
	ID int
	EmployeeId int
	Timesheets []Timesheet
	CreatedAt time.Time
}