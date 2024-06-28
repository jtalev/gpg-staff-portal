package validation_test

import (
	"testing"
	"time"

	"github.com/jtalev/gpg-staff-portal/pkg/types"
	"github.com/jtalev/gpg-staff-portal/pkg/validation"
)

func TestValidTimesheet(t *testing.T) {
	test := types.Result{
		IsValid: true, Error: "",
	}

	timesheet := types.Timesheet{
		Id: 0,
		EmployeeId: 1928374,
		StartTime: time.Date(2024, time.Now().Month(), time.Now().Day(), 7, 0, 0, 0, time.Local),
		EndTime: time.Date(2024, time.Now().Month(), time.Now().Day(), 15, 0, 0, 0, time.Local),
		Date: time.Date(2024, time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local),
		CreatedAt: time.Now(),
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

}