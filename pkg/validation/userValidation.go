package validation

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/jtalev/gpg-staff-portal/pkg/types"
)

func ValidateEmployeeId(id int) types.Result {
	result := types.Result{IsValid: true, Error: ""}
	str := strconv.Itoa(id)

	if len(str) != 7 {
		result = types.Result{IsValid: false, Error: "Incorrect ID length"}
	}

	return result
}

func ValidateFirstName(name string) types.Result {
	result := types.Result{IsValid: true, Error: ""}
	name = strings.Trim(name, " ")

	if len(name) > 25 {
		result = types.Result{IsValid: false, Error: "First name to long"}
	}
	for _, c := range name {
		if unicode.IsDigit(c) {
			result = types.Result{IsValid: false, Error: "First name should not contain digits"}
		}
	}

	return result
}

func ValidateLastName(name string) types.Result {
	result := types.Result{IsValid: true, Error: ""}
	name = strings.Trim(name, " ")

	if len(name) > 25 {
		result = types.Result{IsValid: false, Error: "Last name to long"}
	}
	for _, c := range name {
		if unicode.IsDigit(c) {
			result = types.Result{IsValid: false, Error: "Last name should not contains digits"}
		}
	}

	return result
}

func ValidateEmail(email string) types.Result {
	result := types.Result{IsValid: true, Error: ""}

	r, _ := regexp.Compile("@")
	if !r.MatchString(email) {
		result = types.Result{IsValid: false, Error: "Not a valid email"}
	}

	r, _ = regexp.Compile(".com")
	if !r.MatchString(email) {
		result = types.Result{IsValid: false, Error: "Not a valid email"}
	}

	return result
}

func ValidatePassword(password string) types.Result {
	result := types.Result{IsValid: true, Error: ""}

	if len(password) < 8 {
		result = types.Result{IsValid: false, Error: "Password must contain 8 or more characters"}
	}

	r, _ := regexp.Compile(`[!@#$%^&*(),.?":{}|<>]`)
	if !r.MatchString(password) {
		result = types.Result{IsValid: false, Error: "Password must contain symbol"}
	}

	r, _ = regexp.Compile(`\d`)
	if !r.MatchString(password) {
		result = types.Result{IsValid: false, Error: "Password must contain 1 or more digit(s)"}
	}

	return result
}

func ValidateUser(user types.User) map[string]types.Result {
	results := map[string]types.Result{}

	results["EmployeeId"] = ValidateEmployeeId(user.EmployeeId)
	results["FirstName"] = ValidateFirstName(user.FirstName)
	results["LastName"] = ValidateLastName(user.LastName)
	results["Email"] = ValidateEmail(user.Email)
	results["PasswordHash"] = ValidatePassword(user.PasswordHash)

	return results
}

func ValidateLogin(email, password string) map[string]types.Result {
	results := map[string]types.Result{}

	results["Email"] = ValidateEmail(email)
	results["PasswordHash"] = ValidatePassword(password)

	return results
}
