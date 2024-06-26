package validation

import (
	"github.com/jtalev/gpg-staff-portal/pkg/entities"
)

func ValidateEmployeeId(id int) Result {
	result := Result{IsSuccessful: true, Msg: ""}
	
	if len(string(rune(id))) != 7 {
		result = Result{IsSuccessful: false, Msg: "Employee id should be 7 digits long"}
	}

	return result
}

func ValidateUser(user entities.User) []Result {
	results := make([]Result, 0)
	
	results = append(results, ValidateEmployeeId(user.EmployeeId))

	return results
}