package persistence

import (
	"fmt"
	"log"
	"strconv"
	"time"
	"github.com/jtalev/gpg-staff-portal-api/entities"
)

func getUserDetails(user entities.User) []string {
	var id = fmt.Sprint(user.ID)
	var employeeId = fmt.Sprint(user.EmployeeId)
	var createdAt = user.CreatedAt.String()
	var details = []string{
		id,
		employeeId,
		user.FirstName,
		user.LastName,
		user.Email,
		user.PasswordHash,
		createdAt,
	}

	log.Println("getUserDetails: ", details)
	return details
}

type UserList struct {
	Users []entities.User
}

func GetAllUsers(isReturnError bool) (*UserList, string) {
	if isReturnError {
		userList := UserList{}
		return &userList, "Failed to retrieve data"
	}
	user := entities.User{
		ID:           1,
		EmployeeId:   5613215135,
		FirstName:    "Rob",
		LastName:     "March",
		Email:        "robbie@bigdog.com",
		PasswordHash: "jksbdkjdbsjsdksd",
		CreatedAt:    time.Now(),
	}
	userList := UserList{
		Users: []entities.User{
			user,
		},
	}
	return &userList, "Successfully fetched users"
}

func CreateUser(user entities.User, isReturnError bool) (*[]entities.User, string) {
	userList, _ := GetAllUsers(false)
	if isReturnError {
		return &userList.Users, "Failed to update user"
	}

	for _, property := range getUserDetails(user) {
		if property == "" {
			return &userList.Users, "User detail/s missing"
		}
	}

	for _, obj := range userList.Users {
		if obj.Email == user.Email {
			return &userList.Users, "User already exists"
		}
	}
	userList.Users = append(userList.Users, user)

	return &userList.Users, "Successfully created users"
}

func ReadUser(userId int) (*entities.User, string) {
	userList, _ := GetAllUsers(false)

	for i := range userList.Users {
		if userList.Users[i].ID == userId {
			return &userList.Users[i], "User found"
		}
	}

	user := entities.User{}
	return &user, "User doesn't exist"
}

func UpdateUser(user entities.User, id string) (*entities.User, string) {
	userList, _ := GetAllUsers(false)
	i := -1
	uId, _ := strconv.Atoi(id)

	for i, item := range userList.Users {
		if item.ID == uId {
			userList.Users[i] = user
			return &userList.Users[i], "User found"
		}
	}
	return &userList.Users[i], "User doesn't exist"
}

func DeleteUser(userId int) (*UserList, string) {
	userList, _ := GetAllUsers(false)
	for i, user := range userList.Users {
		if userId == user.ID {
			userList.Users = append(userList.Users[:i], userList.Users[i+1:]...)
			return userList, "User successfully deleted"
		}
	}
	return userList, "User doesn't exist"
}
