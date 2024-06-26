package validation

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/jtalev/gpg-staff-portal/pkg/entities"
)

func ValidateEmployeeId(id int) Result {
	result := Result{IsSuccessful: true, Msg: ""}
	str := strconv.Itoa(id)
	
	if len(str) != 7 {
		result = Result{IsSuccessful: false, Msg: "employee id should be 7 digits long"}
	}

	return result
}

func ValidateFirstName(name string) Result {
	result := Result{IsSuccessful: true, Msg: ""}
	name = strings.Trim(name, " ")

	if len(name) > 25 {
		result = Result{IsSuccessful: false, Msg: "first name to long"}
	}
	for _, c := range name {
		if unicode.IsDigit(c) {
			result = Result{IsSuccessful: false, Msg: "fist name contains digits"}
		}
	}

	return result
}

func ValidateUser(user entities.User) map[string]Result {
	results := map[string]Result{}
	
	results["Employee Id"] = ValidateEmployeeId(user.EmployeeId)

	return results
}