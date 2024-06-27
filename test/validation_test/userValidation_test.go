package validation_test

import (
	"testing"
	"time"

	"github.com/jtalev/gpg-staff-portal/pkg/entities"
	"github.com/jtalev/gpg-staff-portal/pkg/validation"
)

func TestValidUser(t *testing.T) {
	test := validation.Result{
		IsSuccessful: true, Msg: "",
	}

	user := entities.User{

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
		if r.IsSuccessful != test.IsSuccessful {
			t.Errorf("IsSuccessful flag not as expected. want=%v, got=%v",
				test.IsSuccessful, r.IsSuccessful)
		}
		if r.Msg != test.Msg {
			t.Errorf("Msg not as expected: want=%s, got=%s",
				test.Msg, r.Msg)
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
		{validation.Result{IsSuccessful: false, Msg: "Incorrect ID length"}, 19203213, "EmployeeId"},
	}

	strngTests := []validateStringTest{
		{
			validation.Result{IsSuccessful: false, Msg: "First name to long"},
			"abcdefghijklmnoqrstuvwxyzabcdefg",
			"FirstName",
		},
		{
			validation.Result{IsSuccessful: false, Msg: "First name should not contain digits"},
			"sl1ddy",
			"FirstName",
		},
		{
			validation.Result{IsSuccessful: false, Msg: "Last name to long"},
			"abcdefghijklmnoqrstuvwxyzabcdefg",
			"LastName",
		},
		{
			validation.Result{IsSuccessful: false, Msg: "Last name should not contains digits"},
			"sl1ddy",
			"LastName",
		},
		{
			validation.Result{IsSuccessful: false, Msg: "Not a valid email"},
			"bigfellaoutlook.com",
			"Email",
		},
		{
			validation.Result{IsSuccessful: false, Msg: "Not a valid email"},
			"bigfella@outlook",
			"Email",
		},
		{
			validation.Result{IsSuccessful: false, Msg: "Password must contain 8 or more characters"},
			"1!",
			"PasswordHash",
		},
		{
			validation.Result{IsSuccessful: false, Msg: "Password must contain symbol"},
			"password1",
			"PasswordHash",
		},
		{
			validation.Result{IsSuccessful: false, Msg: "Password must contain 1 or more digit(s)"},
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
	e := validation.Result{IsSuccessful: true, Msg: ""}
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
