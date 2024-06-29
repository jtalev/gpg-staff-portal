package validation_test

import (
	"testing"
	"time"

	"github.com/jtalev/gpg-staff-portal/pkg/types"
	"github.com/jtalev/gpg-staff-portal/pkg/validation"
)

func TestValidTimesheet(t *testing.T) {
	test := types.ValidationResult{
		IsValid: true, Error: "",
	}

	timesheet := types.Timesheet{
		Id:         0,
		EmployeeId: 1928374,
		StartTime:  time.Date(2024, time.Now().Month(), time.Now().Day(), 7, 0, 0, 0, time.Local),
		EndTime:    time.Date(2024, time.Now().Month(), time.Now().Day(), 15, 0, 0, 0, time.Local),
		Date:       time.Date(2024, time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local),
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	}

	results := validation.ValidateTimesheet(timesheet)

	for _, r := range results {
		if r.IsValid != test.IsValid {
			t.Errorf("IsValid flag not as expected. want=%v, got=%v", test.IsValid, r.IsValid)
		}
		if r.Error != test.Error {
			t.Errorf("Error message not as expected. want=%v, got=%v", test.Error, r.Error)
		}
	}
}

func TestInvalidTimesheetField(t *testing.T) {
	intTests := []types.ValidateIntTest{
		{
			Expected: types.ValidationResult{IsValid: false, Error: "Incorrect ID length"},
			Value:    19283743,
			Field:    "EmployeeId",
		},
		{
			Expected: types.ValidationResult{IsValid: false, Error: "Incorrect ID length"},
			Value:    192837,
			Field:    "EmployeeId",
		},
	}

	timeTests := []types.ValidateTimeTest{
		{
			Expected: types.ValidationResult{IsValid: false, Error: "Start time must be set"},
			Start:    time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			End:      time.Date(1, time.January, 1, 1, 0, 0, 0, time.UTC),
			Date:     time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			Field:    "StartTime",
		},
		{
			Expected: types.ValidationResult{IsValid: false, Error: "Start time must be before end time"},
			Start:    time.Date(1, time.January, 1, 1, 0, 0, 0, time.UTC),
			End:      time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			Date:     time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			Field:    "StartTime",
		},
		{
			Expected: types.ValidationResult{IsValid: false, Error: "Start time must have same date as timesheet"},
			Start:    time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			End:      time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			Date:     time.Now(),
			Field:    "StartTime",
		},
		{
			Expected: types.ValidationResult{IsValid: false, Error: "End time must be set"},
			Start:    time.Date(1, time.January, 1, -1, 0, 0, 0, time.UTC),
			End:      time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			Date:     time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			Field:    "EndTime",
		},
		{
			Expected: types.ValidationResult{IsValid: false, Error: "End time must be after start time"},
			Start:    time.Date(1, time.January, 1, 1, 0, 0, 0, time.UTC),
			End:      time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			Date:     time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			Field:    "EndTime",
		},
		{
			Expected: types.ValidationResult{IsValid: false, Error: "End time must have same date as timesheet"},
			Start:    time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			End:      time.Date(1, time.January, 1, 1, 0, 0, 0, time.UTC),
			Date:     time.Now(),
			Field:    "EndTime",
		},
		{
			Expected: types.ValidationResult{IsValid: false, Error: "Date must be set"},
			Start:    time.Now(),
			End:      time.Now(),
			Date:     time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			Field:    "Date",
		},
		{
			Expected: types.ValidationResult{IsValid: false, Error: "Date must not be a future date"},
			Start:    time.Now(),
			End:      time.Now(),
			Date:     time.Now().AddDate(0, 0, 1),
			Field:    "Date",
		},
	}

	for _, tt := range intTests {
		if tt.Field == "EmployeeId" {
			r := validation.ValidateEmployeeId(tt.Value)
			if r != tt.Expected {
				t.Errorf("EmployeeId: want=%v, got=%v", tt.Expected, r)
			}
		}
	}

	for _, tt := range timeTests {
		if tt.Field == "StartTime" {
			r := validation.ValidateStartTime(tt.Start, tt.End, tt.Date)
			if r != tt.Expected {
				t.Errorf("StartTime: want=%v, got=%v", tt.Expected, r)
			}
		}
		if tt.Field == "EndTime" {
			r := validation.ValidateEndTime(tt.Start, tt.End, tt.Date)
			if r != tt.Expected {
				t.Errorf("EndTime: want=%v, got=%v", tt.Expected, r)
			}
		}
		if tt.Field == "Date" {
			r := validation.ValidateDate(tt.Date)
			if r != tt.Expected {
				t.Errorf("Date: want=%v, got=%v", tt.Expected, r)
			}
		}
	}
}
