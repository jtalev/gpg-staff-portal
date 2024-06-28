package mockdb

import (
	"time"

	"github.com/jtalev/gpg-staff-portal/pkg/types"
)

var users = []types.User{
	{
		Uid:          0,
		EmployeeId:   000000,
		FirstName:    "Robbie",
		LastName:     "Marsh",
		Email:        "robbieLad@outlook.com",
		PasswordHash: "blsbdjksd65452",
		CreatedAt:    time.Now(),
		ModifiedAt:   time.Now(),
	},
	{
		Uid:          1,
		EmployeeId:   000001,
		FirstName:    "Ronnie",
		LastName:     "March",
		Email:        "ronDawg@outlook.com",
		PasswordHash: "kjbslkbsdk6+23",
		CreatedAt:    time.Now(),
		ModifiedAt:   time.Now(),
	},
}

func GetUserData() ([]types.User, error) {
	return users, nil
}

func GetUserById(id int) (*types.User, string) {
	for _, user := range users {
		if user.Uid == id {
			return &user, "Successfully fetched user"
		}
	}

	return nil, "User doesn't exist"
}

func GetUserByEmployeeId(employeeId int) (*types.User, string) {
	for _, user := range users {
		if user.EmployeeId == employeeId {
			return &user, "Successfully fetched user"
		}
	}

	return nil, "User doesn't exist"
}

func CreateUser(user types.User) (*[]types.User, string) {
	for _, item := range users {
		if user.EmployeeId == item.EmployeeId {
			return nil, "Employee ID already exists"
		}
		if user.Email == item.Email {
			return nil, "Email already exists"
		}
	}

	users = append(users, user)
	return &users, "Successfully created user"
}

func UpdateUser(user types.User, id int) (*types.User, string) {
	for i, item := range users {
		if id == item.Uid {
			users[i] = user
			return &user, "User successfully updated"
		}
	}
	return nil, "User doesn't exist"
}

func DeleteUser(id int) (*[]types.User, string) {
	for i, user := range users {
		if user.Uid == id {
			before := users[:i]
			after := users[i+1:]
			users = append(before, after...)
			return &users, "User successfully deleted"
		}
	}
	return nil, "User doesn't exist"
}
