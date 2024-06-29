package mockdb

import (
	"fmt"
	"time"

	"github.com/jtalev/gpg-staff-portal/pkg/types"
)

type TimesheetRepo struct {}

var timesheets = []types.Timesheet{
	{
		Id:         0,
		EmployeeId: 1234567,
		StartTime:  time.Now(),
		EndTime:    time.Now(),
		Date:       time.Date(2024, time.June, 17, 8, 0, 0, 0, time.Now().Location()),
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	},
	{
		Id:         1,
		EmployeeId: 1234567,
		StartTime:  time.Now(),
		EndTime:    time.Now(),
		Date:       time.Date(2024, time.June, 17, 8, 0, 0, 0, time.Now().Location()),
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	},
	{
		Id:         2,
		EmployeeId: 1234568,
		StartTime:  time.Now(),
		EndTime:    time.Now(),
		Date:       time.Date(2024, time.June, 17, 8, 0, 0, 0, time.Now().Location()),
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	},
}

func (t *TimesheetRepo) GetAllTimesheets() (ts *[]types.Timesheet, msg string) {
	return &timesheets, "Successfully fetched timesheets"
}

func (t *TimesheetRepo) GetTimesheetsByEmployeeId(empId int) (ts *[]types.Timesheet, msg string) {
	tsList := make([]types.Timesheet, 0)
	for _, t := range timesheets {
		if t.EmployeeId == empId {
			tsList = append(tsList, t)
		}
	}

	if len(tsList) == 0 {
		return nil, "User not assigned to any timesheets"
	}
	return &tsList, "Successfully fetched users timesheets"
}

func (t *TimesheetRepo) CreateTimesheet(timesheet types.Timesheet) (ts *types.Timesheet, msg string) {
	tsSlice, _ := t.GetTimesheetsByEmployeeId(timesheet.EmployeeId)
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

func (t *TimesheetRepo) UpdateTimesheet(timesheet types.Timesheet, id int) (ts *types.Timesheet, msg string) {
	for i, ts := range timesheets {
		if ts.Id == id {
			timesheets[i] = timesheet
			return &timesheet, "Successfully updated timesheet"
		}
	}
	return nil, "No timesheet with given id"
}

func (t *TimesheetRepo) DeleteTimesheet(id int) (ts *[]types.Timesheet, msg string) {
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

func (t *TimesheetRepo) GetTimesheetWithinRange(
	id int, start, end time.Time,
	)(
		ts *[]types.Timesheet, msg string,
){
	tsById, _ := t.GetTimesheetsByEmployeeId(id)
	fmt.Println(len(*tsById))

	tss := make([]types.Timesheet, 0)
	start = start.AddDate(0, 0, -1)
	end = end.AddDate(0, 0, 1)

	for _, t := range *tsById {
		if t.Date.After(start) && t.Date.Before(end) {
			tss = append(tss, t)
		}
	}

	return &tss, "Successfully fetched timesheets"
}
