package mockdb

import (
	"time"

	"github.com/jtalev/gpg-staff-portal/pkg/entities"
)

var timesheets = []entities.Timesheet{
	{
		Id:         0,
		EmployeeId: 000000,
		StartTime:  time.Now(),
		EndTime:    time.Now(),
		Date:       time.Date(2024, time.June, 17, 8, 0, 0, 0, time.Now().Location()),
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	},
	{
		Id:         1,
		EmployeeId: 000000,
		StartTime:  time.Now(),
		EndTime:    time.Now(),
		Date:       time.Date(2024, time.June, 17, 8, 0, 0, 0, time.Now().Location()),
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	},
	{
		Id:         2,
		EmployeeId: 000001,
		StartTime:  time.Now(),
		EndTime:    time.Now(),
		Date:       time.Date(2024, time.June, 17, 8, 0, 0, 0, time.Now().Location()),
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	},
}

func GetAllTimesheets() (t *[]entities.Timesheet, msg string) {
	return &timesheets, "Successfully fetched timesheets"
}

func GetTimesheetsByEmployeeId(empId int) (t *[]entities.Timesheet, msg string) {
	ts := make([]entities.Timesheet, 0)
	for _, t := range timesheets {
		if t.EmployeeId == empId {
			ts = append(ts, t)
		}
	}

	if len(ts) == 0 {
		return nil, "User not assigned to any timesheets"
	}
	return &ts, "Successfully fetched users timesheets"
}

func CreateTimesheet(timesheet entities.Timesheet) (t *entities.Timesheet, msg string) {
	tsSlice, _ := GetTimesheetsByEmployeeId(timesheet.EmployeeId)
	for _, ts := range *tsSlice {
		if ts.Id == timesheet.Id {
			return nil, "Timesheet already exists"
		}
		if ts.Date == timesheet.Date {
			return nil, "Timesheet already created for this date, use update"
		}
	}
	timesheets = append(timesheets, timesheet)
	return &timesheet, "Timesheet successfully created"
}

func UpdateTimesheet(timesheet entities.Timesheet, id int) (t *entities.Timesheet, msg string) {
	for i, ts := range timesheets {
		if ts.Id == id {
			timesheets[i] = timesheet
			return &timesheet, "Successfully updated timesheet"
		}
	}
	return nil, "No timesheet with given id"
}

func DeleteTimesheet(id int) (t *[]entities.Timesheet, msg string) {
	for i, ts := range timesheets {
		if ts.Id == id {
			before := timesheets[:i]
			after := timesheets[i+1:]
			timesheets = append(before, after...)
			return &timesheets, "Successfully deleted timesheet"
		}
	}
	return nil, "Timesheet with given id doesn't exist"
}
