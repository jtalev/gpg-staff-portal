package validation

import (
	"time"

	"github.com/jtalev/gpg-staff-portal/pkg/types"
)

func ValidateStartDate(s, e time.Time) types.ValidationResult {
	result := types.ValidationResult{IsValid: true, Error: ""}

	if s.IsZero() {
		result = types.ValidationResult{
			IsValid: false, 
			Error: "Start date must be set",
		}
	}
	if s.After(e) {
		result = types.ValidationResult{
			IsValid: false,
			Error: "Start date must be before end date",
		}
	}

	return result
}

func ValidateEndDate(s, e time.Time) types.ValidationResult {
	result := types.ValidationResult{IsValid: true, Error: ""}

	if e.IsZero() {
		result = types.ValidationResult{
			IsValid: false, 
			Error: "End date must be set",
		}
	}
	if s.After(e) {
		result = types.ValidationResult{
			IsValid: false,
			Error: "End date must be after start date",
		}
	}

	return result
}

func ValidateLeaveRequest(l types.LeaveRequest) map[string]types.ValidationResult {
	results := map[string]types.ValidationResult{}

	results["StartDate"] = ValidateStartDate(l.StartDate, l.EndDate)
	results["EndDate"] = ValidateEndDate(l.StartDate, l.EndDate)

	return results
}