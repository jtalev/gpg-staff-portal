package validation_test

import (
	"testing"
	"time"

	"github.com/jtalev/gpg-staff-portal/pkg/types"
	"github.com/jtalev/gpg-staff-portal/pkg/validation"
)

func TestValidLeaveRequest(t *testing.T) {
	test := types.ValidationResult{
		IsValid: true, Error: "",
	}

	start := time.Now()
	end := start.AddDate(0, 0, 10)
	l := types.LeaveRequest{
		Id: 0,
		EmployeeId: 1532645,
		Type: "annual",
		StartDate: start,
		EndDate: end,
		TotalDays: 10,
		IsApproved: true,
		CreatedAt: time.Now(),
		ModifiedAt: time.Now(),
	}

	results := validation.ValidateLeaveRequest(l)

	for _, r := range results {
		if r.IsValid != test.IsValid {
			t.Errorf("IsValid flag not as expected. want=%v, got=%v", test.IsValid, r.IsValid)
		}
		if r.Error != test.Error {
			t.Errorf("Error message not as expected. want=%v, got=%v", test.Error, r.Error)
		}
	}
}

func TestInvalidLeaveRequestField(t *testing.T) {
	timeTests := []types.ValidateTimeTest{
		{
			Expected: types.ValidationResult{
				IsValid: false, Error: "Start date must be set",
			},
			Start: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			End: time.Now(),
			Date: time.Now(),
			Field: "StartDate",
		},
		{
			Expected: types.ValidationResult{
				IsValid: false, Error: "Start date must be before end date",
			},
			Start: time.Now(),
			End: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			Date: time.Now(),
			Field: "StartDate",
		},
		{
			Expected: types.ValidationResult{
				IsValid: false, Error: "End date must be set",
			},
			Start: time.Date(1, time.January, -1, 0, 0, 0, 0, time.UTC),
			End: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			Date: time.Now(),
			Field: "EndDate",
		},
		{
			Expected: types.ValidationResult{
				IsValid: false, Error: "End date must be after start date",
			},
			Start: time.Now(),
			End: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			Date: time.Now(),
			Field: "EndDate",
		},
	}

	for _, tt := range timeTests {
		if tt.Field == "StartDate" {
			r := validation.ValidateStartDate(tt.Start, tt.End)
			if r != tt.Expected {
				t.Errorf("StartDate: want=%v, got=%v", tt.Expected, r)
			}
		}
		if tt.Field == "EndDate" {
			r := validation.ValidateEndDate(tt.Start, tt.End)
			if r != tt.Expected {
				t.Errorf("EndDate: want=%v, got=%v", tt.Expected, r)
			}
		}
	}
}
