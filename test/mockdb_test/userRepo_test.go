package mockdb_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/jtalev/gpg-staff-portal/internal/entities"
	"github.com/jtalev/gpg-staff-portal/internal/mockdb"
)

func TestGetUserData(t *testing.T) {
	robbie := entities.User{
		Uid:          0,
		EmployeeId:   000000,
		FirstName:    "Robbie",
		LastName:     "Marsh",
		Email:        "robbieLad@outlook.com",
		PasswordHash: "blsbdjksd65452",
		CreatedAt:    time.Now(),
		ModifiedAt:   time.Now(),
	}
	ronnie := entities.User{
		Uid:          1,
		EmployeeId:   000001,
		FirstName:    "Ronnie",
		LastName:     "March",
		Email:        "ronDawg@outlook.com",
		PasswordHash: "kjbslkbsdk6+23",
		CreatedAt:    time.Now(),
		ModifiedAt:   time.Now(),
	}
	data := []entities.User{
		robbie,
		ronnie,
	}

	got, _ := mockdb.GetUserData()
	if !reflect.DeepEqual(got, data) {
		t.Errorf("GetUserData() = %v; want %v", got, data)
	}
}

func TestGetUserById(t *testing.T) {
	robbie := entities.User{
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
	got, _ := mockdb.GetUserById(0)
	if !reflect.DeepEqual(got, &robbie) {
		t.Errorf("GetUserById(0) = %v; want %v", got, &robbie)
	}

	// user id not in db
	got, err := mockdb.GetUserById(999)
	if err == "Successfully fetched user" {
		t.Errorf("GetUserById(999) error = %v; want User doesn't exist", err)
	}
	if got != nil {
		t.Errorf("GetUserById(999) = %v; want nil", got)
	}
}

func TestGetUserByEmployeeId(t *testing.T) {
	robbie := entities.User{
		Uid:          0,
		EmployeeId:   000000,
		FirstName:    "Robbie",
		LastName:     "Marsh",
		Email:        "robbieLad@outlook.com",
		PasswordHash: "blsbdjksd65452",
		CreatedAt:    time.Now(),
		ModifiedAt:   time.Now(),
	}

	// employee id in db
	got, _ := mockdb.GetUserByEmployeeId(000000)
	if !reflect.DeepEqual(got, &robbie) {
		t.Errorf("GetUserByEmployeeId(000000) = %v; want %v", got, &robbie)
	}

	// employee id not in db
	got, err := mockdb.GetUserByEmployeeId(521213)
	if err == "Successfully fetched user" {
		t.Errorf("GetUserByEmployeeId(000000) error = %v; want User doesn't exist", err)
	}
	if got != nil {
		t.Errorf("GetUserByEmployeeId(000000) = %v; want nil", got)
	}
}
