package mockdb_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/jtalev/gpg-staff-portal/pkg/mockdb"
	"github.com/jtalev/gpg-staff-portal/pkg/types"
)

var u = mockdb.UserRepo{}

func TestGetUserData(t *testing.T) {
	robbie := types.User{
		Uid:          0,
		EmployeeId:   000000,
		FirstName:    "Robbie",
		LastName:     "Marsh",
		Email:        "robbieLad@outlook.com",
		PasswordHash: "blsbdjksd65452",
		CreatedAt:    time.Now(),
		ModifiedAt:   time.Now(),
	}
	ronnie := types.User{
		Uid:          1,
		EmployeeId:   000001,
		FirstName:    "Ronnie",
		LastName:     "March",
		Email:        "ronDawg@outlook.com",
		PasswordHash: "kjbslkbsdk6+23",
		CreatedAt:    time.Now(),
		ModifiedAt:   time.Now(),
	}
	data := []types.User{
		robbie,
		ronnie,
	}

	got, _ := u.GetUserData()
	if !reflect.DeepEqual(got, data) {
		t.Errorf("GetUserData() = %v; want %v", got, data)
	}
}

func TestGetUserById(t *testing.T) {
	robbie := types.User{
		Uid:          0,
		EmployeeId:   000000,
		FirstName:    "Robbie",
		LastName:     "Marsh",
		Email:        "robbieLad@outlook.com",
		PasswordHash: "blsbdjksd65452",
		CreatedAt:    time.Now(),
		ModifiedAt:   time.Now(),
	}

	// user id in db
	got, _ := u.GetUserById(0)
	if !reflect.DeepEqual(got, &robbie) {
		t.Errorf("GetUserById(0) = %v; want %v", got, &robbie)
	}

	// user id not in db
	got, err := u.GetUserById(999)
	if err == "Successfully fetched user" {
		t.Errorf("GetUserById(999) error = %v; want User doesn't exist", err)
	}
	if got != nil {
		t.Errorf("GetUserById(999) = %v; want nil", got)
	}
}

func TestGetUserByEmployeeId(t *testing.T) {
	robbie := types.User{
		Uid:          0,
		EmployeeId:   000000,
		FirstName:    "Robbie",
		LastName:     "Marsh",
		Email:        "robbieLad@outlook.com",
		PasswordHash: "blsbdjksd65452",
		CreatedAt:    time.Now(),
		ModifiedAt:   time.Now(),
	}

	got, _ := u.GetUserByEmployeeId(000000)
	if !reflect.DeepEqual(got, &robbie) {
		t.Errorf("GetUserByEmployeeId(000000) = %v; want %v", got, &robbie)
	}

	got, err := u.GetUserByEmployeeId(521213)
	if err == "Successfully fetched user" {
		t.Errorf("GetUserByEmployeeId(000000) error = %v; want User doesn't exist", err)
	}
	if got != nil {
		t.Errorf("GetUserByEmployeeId(000000) = %v; want nil", got)
	}
}
