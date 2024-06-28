package validation_test

import (
	"testing"
	"time"

	"github.com/jtalev/gpg-staff-portal/pkg/types"
	"github.com/jtalev/gpg-staff-portal/pkg/validation"
)

func TestValidUser(t *testing.T) {
	test := validation.Result{
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
			t.Errorf("IsSuccessful flag not as expected. want=%v, got=%v",
				test.IsValid, r.IsValid)
		}
		if r.Error != test.Error {
			t.Errorf("Msg not as expected: want=%s, got=%s",
				test.Error, r.Error)
		}
	}
}

type validateIntTest struct {
	expected validation.Result
	value    int
	field    string
}

type validateStringTest struct {
	expected validation.Result
	value    string
	field    string
}

func TestFailedValidation(t *testing.T) {
	intTests := []validateIntTest{
		{validation.Result{IsValid: false, Error: "Incorrect ID length"}, 19203213, "EmployeeId"},
	}

	strngTests := []validateStringTest{
		{
			validation.Result{IsValid: false, Error: "First name to long"},
			"abcdefghijklmnoqrstuvwxyzabcdefg",
			"FirstName",
		},
		{
			validation.Result{IsValid: false, Error: "First name should not contain digits"},
			"sl1ddy",
			"FirstName",
		},
		{
			validation.Result{IsValid: false, Error: "Last name to long"},
			"abcdefghijklmnoqrstuvwxyzabcdefg",
			"LastName",
		},
		{
			validation.Result{IsValid: false, Error: "Last name should not contains digits"},
			"sl1ddy",
			"LastName",
		},
		{
			validation.Result{IsValid: false, Error: "Not a valid email"},
			"bigfellaoutlook.com",
			"Email",
		},
		{
			validation.Result{IsValid: false, Error: "Not a valid email"},
			"bigfella@outlook",
			"Email",
		},
		{
			validation.Result{IsValid: false, Error: "Password must contain 8 or more characters"},
			"1!",
			"PasswordHash",
		},
		{
			validation.Result{IsValid: false, Error: "Password must contain symbol"},
			"password1",
			"PasswordHash",
		},
		{
			validation.Result{IsValid: false, Error: "Password must contain 1 or more digit(s)"},
			"password!",
			"PasswordHash",
		},
	}

	for _, tt := range intTests {
		if tt.field == "EmployeeId" {
			r := validation.ValidateEmployeeId(tt.value)
			if r != tt.expected {
				t.Errorf("EmployeeId: want=%v, got=%v", tt.expected, r)
			}
		}
	}

	for _, tt := range strngTests {
		if tt.field == "FirstName" {
			r := validation.ValidateFirstName(tt.value)
			if r != tt.expected {
				t.Errorf("FirstName: want=%v, got=%v", tt.expected, r)
			}
		}
		if tt.field == "LastName" {
			r := validation.ValidateLastName(tt.value)
			if r != tt.expected {
				t.Errorf("LastName: want=%v, got=%v", tt.expected, r)
			}
		}
		if tt.field == "Email" {
			r := validation.ValidateEmail(tt.value)
			if r != tt.expected {
				t.Errorf("Email: want=%v, got=%v", tt.expected, r)
			}
		}
		if tt.field == "PasswordHash" {
			r := validation.ValidatePassword(tt.value)
			if r != tt.expected {
				t.Errorf("Password: want=%v, got=%v", tt.expected, r)
			}
		}
	}
}

type Login struct {
	email    string
	password string
}

func TestValidLogin(t *testing.T) {
	e := validation.Result{IsValid: true, Error: ""}
	l := Login{"bigfella@outlook.com", "password1!"}
	r := validation.ValidateLogin(l.email, l.password)

	if r["Email"] != e {
		t.Errorf("Email: want=%v, got=%v", e, r["Email"])
	}
	if r["PasswordHash"] != e {
		t.Errorf("PasswordHash: want=%v, got=%v", e, r["Email"])
	}
}

// TODO: test for incorrect login field input
