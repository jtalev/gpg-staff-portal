package validation_test

import (
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
		EmployeeId:   000000,
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