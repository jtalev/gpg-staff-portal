package entities

import "time"

type Timesheet struct {
	Id         int
	EmployeeId int
	StartTime  time.Time
	EndTime    time.Time
	Date  	   time.Time
	CreatedAt  time.Time
	ModifiedAt time.Time
}