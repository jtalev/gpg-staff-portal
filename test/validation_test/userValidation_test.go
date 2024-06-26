package validation_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/jtalev/gpg-staff-portal/pkg/entities"
	"github.com/jtalev/gpg-staff-portal/pkg/validation"
)

func TestValidateUser(t *testing.T) {
	test := validation.Result{
		IsSuccessful: true, Msg: "",
	}

	user := entities.User{
		
		Uid:          0,
		EmployeeId:   1920321,
		FirstName:    "Robbie",
		LastName:     "Marsh",
		Email:        "robbieLad@outlook.com",
		PasswordHash: "blsbdjksd65452",
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

func TestFailedValidation(t *testing.T) {
	expected := map[string]validation.Result{
		"Employee Id": {IsSuccessful: false, Msg: "employee id should be 7 digits long"},
	}

	user := entities.User{
		
		Uid:          0,
		EmployeeId:   19203213,
		FirstName:    "Robbie",
		LastName:     "Marsh",
		Email:        "robbieLad@outlook.com",
		PasswordHash: "blsbdjksd65452",
		CreatedAt:    time.Now(),
		ModifiedAt:   time.Now(),
		
	}

	result := validation.ValidateUser(user)

	for i := 0; i < len(result); i++{
		if result["Employee Id"] != expected["Employee Id"] {
			t.Errorf("Employee Id validation failed. want=%v, got %v", 
					expected["Employee Id"], result["Employee Id"])
		}
	}
}