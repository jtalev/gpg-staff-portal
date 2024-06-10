package entities

type Paycycle struct {
	Id int
	EmployeeId int
	Timesheets []Timesheet
}