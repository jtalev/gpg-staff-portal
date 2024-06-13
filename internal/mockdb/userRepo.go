package mockdb

import (
	"time"

	"github.com/jtalev/gpg-staff-portal/internal/entities"
)

func GetUserData() ([]entities.User, string) {

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

	return data, "Successfully fetched users"
}

func GetUserById(id int) (*entities.User, string) {

	data, _ := GetUserData()
	for _, user := range data {
		if user.Uid == id {
			return &user, "Successfully fetched user"
		}
	}

	return nil, "User doesn't exist"
}

func GetUserByEmployeeId(employeeId int) (*entities.User, string) {

	data, _ := GetUserData()
	for _, user := range data {
		if user.EmployeeId == employeeId {
			return &user, "Successfully fetched user"
		}
	}

	return nil, "User doesn't exist"
}

func CreateUser(user entities.User) (*[]entities.User, string) {

	data, _ := GetUserData()
	for _, item := range data {
		if user.EmployeeId == item.EmployeeId {
			return nil, "Employee ID already exists"
		}
		if user.Email == item.Email {
			return nil, "Email already exists"
		}
	}

	data = append(data, user)
	return &data, "Successfully created user"
}

func UpdateUser(user entities.User, id int) (*entities.User, string) {

	data, _ := GetUserData()
	for i, item := range data {
		if id == item.Uid {
			data[i] = user
			return &user, "User successfully updated"
		}
	}
	return nil, "User doesn't exist"
}

func DeleteUser(id int) (*[]entities.User, string) {

	data, _ := GetUserData()
	for i, user := range data {
		if user.Uid == id {
			before := data[:i]
			after := data[i+1:]
			data = append(before, after...)
			return &data, "User successfully deleted"
		}
	}
	return nil, "User doesn't exist"
}
