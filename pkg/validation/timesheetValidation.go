package validation

import (
	"time"

	"github.com/jtalev/gpg-staff-portal/pkg/types"
)

func ValidateStartTime(start, end, date time.Time) types.ValidationResult {
	result := types.ValidationResult{IsValid: true, Error: ""}

	if start.IsZero() {
		result = types.ValidationResult{IsValid: false, Error: "Start time must be set"}
	}
	if !start.Before(end) {
		result = types.ValidationResult{IsValid: false, Error: "Start time must be before end time"}
	}
	if start.Year() != date.Year() || start.YearDay() != date.YearDay() {
		result = types.ValidationResult{IsValid: false, Error: "Start time must have same date as timesheet"}
	}

	return result
}

func ValidateEndTime(start, end, date time.Time) types.ValidationResult {
	result := types.ValidationResult{IsValid: true, Error: ""}

	if end.IsZero() {
		result = types.ValidationResult{IsValid: false, Error: "End time must be set"}
	}
	if !end.After(start) {
		result = types.ValidationResult{IsValid: false, Error: "End time must be after start time"}
	}
	if end.Year() != date.Year() || end.YearDay() != date.YearDay() {
		result = types.ValidationResult{IsValid: false, Error: "End time must have same date as timesheet"}
	}

	return result
}

func ValidateDate(date time.Time) types.ValidationResult {
	result := types.ValidationResult{IsValid: true, Error: ""}
	now := time.Now()

	if date.IsZero() {
		result = types.ValidationResult{IsValid: false, Error: "Date must be set"}
	}
	if date.After(now) {
		result = types.ValidationResult{IsValid: false, Error: "Date must not be a future date"}
	}

	return result
}

func ValidateTimesheet(t types.Timesheet) map[string]types.ValidationResult {
	results := map[string]types.ValidationResult{}

	results["EmployeeId"] = ValidateEmployeeId(t.EmployeeId)
	results["StartTime"] = ValidateStartTime(t.StartTime, t.EndTime, t.Date)
	results["Endtime"] = ValidateEndTime(t.StartTime, t.EndTime, t.Date)
	results["Date"] = ValidateDate(t.Date)

	return results
}
