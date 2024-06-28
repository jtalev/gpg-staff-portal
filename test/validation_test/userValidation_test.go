package validation_test

import (
	"testing"
	"time"

	"github.com/jtalev/gpg-staff-portal/pkg/validation"
	"github.com/jtalev/gpg-staff-portal/pkg/types"
)

func TestValidUser(t *testing.T) {
	test := types.Result{
		IsValid: true, Error: "",
	}

	user := types.User{

		Uid:          0,
		EmployeeId:   1920321,
		FirstName:    "Robbie",
		LastName:     "Marsh",
		Email:        "robbieLad@outlook.com",
		PasswordHash: "blsbdjksd65452!",
		CreatedAt:    time.Now(),
		ModifiedAt:   time.Now(),
	}

	result := validation.ValidateUser(user)

	for _, r := range result {
		if r.IsValid != test.IsValid {
			t.Errorf("IsValid flag not as Expected. want=%v, got=%v",
				test.IsValid, r.IsValid)
		}
		if r.Error != test.Error {
			t.Errorf("Error not as Expected: want=%s, got=%s",
				test.Error, r.Error)
		}
	}
}

func TestInvalidUserField(t *testing.T) {
	intTests := []types.ValidateIntTest{
		{
			Expected: types.Result{IsValid: false, Error: "Incorrect ID length"},
			Value: 19203213, 
			Field: "EmployeeId",
		},
	}

	strngTests := []types.ValidateStringTest{
		{
			Expected: types.Result{IsValid: false, Error: "First name to long"},
			Value: "abcdefghijklmnoqrstuvwxyzabcdefg",
			Field: "FirstName",
		},
		{
			Expected: types.Result{IsValid: false, Error: "First name should not contain digits"},
			Value: "sl1ddy",
			Field: "FirstName",
		},
		{
			Expected: types.Result{IsValid: false, Error: "Last name to long"},
			Value: "abcdefghijklmnoqrstuvwxyzabcdefg",
			Field: "LastName",
		},
		{
			Expected: types.Result{IsValid: false, Error: "Last name should not contains digits"},
			Value: "sl1ddy",
			Field: "LastName",
		},
		{
			Expected: types.Result{IsValid: false, Error: "Not a valid email"},
			Value: "bigfellaoutlook.com",
			Field: "Email",
		},
		{
			Expected: types.Result{IsValid: false, Error: "Not a valid email"},
			Value: "bigfella@outlook",
			Field: "Email",
		},
		{
			Expected: types.Result{IsValid: false, Error: "Password must contain 8 or more characters"},
			Value: "1!",
			Field: "PasswordHash",
		},
		{
			Expected: types.Result{IsValid: false, Error: "Password must contain symbol"},
			Value: "password1",
			Field: "PasswordHash",
		},
		{
			Expected: types.Result{IsValid: false, Error: "Password must contain 1 or more digit(s)"},
			Value: "password!",
			Field: "PasswordHash",
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

	for _, tt := range strngTests {
		if tt.Field == "FirstName" {
			r := validation.ValidateFirstName(tt.Value)
			if r != tt.Expected {
				t.Errorf("FirstName: want=%v, got=%v", tt.Expected, r)
			}
		}
		if tt.Field == "LastName" {
			r := validation.ValidateLastName(tt.Value)
			if r != tt.Expected {
				t.Errorf("LastName: want=%v, got=%v", tt.Expected, r)
			}
		}
		if tt.Field == "Email" {
			r := validation.ValidateEmail(tt.Value)
			if r != tt.Expected {
				t.Errorf("Email: want=%v, got=%v", tt.Expected, r)
			}
		}
		if tt.Field == "PasswordHash" {
			r := validation.ValidatePassword(tt.Value)
			if r != tt.Expected {
				t.Errorf("Password: want=%v, got=%v", tt.Expected, r)
			}
		}
	}
}

func TestValidLogin(t *testing.T) {
	e := types.Result{IsValid: true, Error: ""}
	l := types.Login{Email: "bigfella@outlook.com", Password: "password1!"}
	r := validation.ValidateLogin(l.Email, l.Password)

	if r["Email"] != e {
		t.Errorf("Email: want=%v, got=%v", e, r["Email"])
	}
	if r["PasswordHash"] != e {
		t.Errorf("PasswordHash: want=%v, got=%v", e, r["Email"])
	}
}