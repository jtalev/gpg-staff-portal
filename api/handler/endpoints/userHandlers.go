package endpoints

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jtalev/gpg-staff-portal/pkg/mockdb"
	"github.com/jtalev/gpg-staff-portal/pkg/types"
)

type UserHandler struct {
	Db mockdb.Db
}

func (u *UserHandler) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, _ := u.Db.UserRepo.GetUserData()

	path := filepath.Join("..", "..", "web", "tmpl", "userList.html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		log.Println("Failed to parse file: ", err)
	}

	err = tmpl.Execute(w, users)
	if err != nil {
		log.Println("Failed to execute template: ", err)
	}
}

func (u *UserHandler) GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		log.Println("id invalid")
		return
	}

	user, msg := u.Db.UserRepo.GetUserById(id)
	encoder := json.NewEncoder(w)
	encoder.Encode(user)
	encoder.Encode(msg)
}

func (u *UserHandler) GetUserByEmployeeIdHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	empIdStr := vars["employeeId"]

	id, err := strconv.Atoi(empIdStr)
	if err != nil {
		http.Error(w, "invalid employee id", http.StatusBadRequest)
		log.Println("id employee invalid")
		return
	}

	user, msg := u.Db.UserRepo.GetUserByEmployeeId(id)
	encoder := json.NewEncoder(w)
	encoder.Encode(user)
	encoder.Encode(msg)
}

func (u *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user types.User

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		log.Println("Error decoding JSON: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user.CreatedAt = time.Now()
	user.ModifiedAt = time.Now()

	log.Println(user)

	response, msg := u.Db.UserRepo.CreateUser(user)

	json.NewEncoder(w).Encode(response)
	json.NewEncoder(w).Encode(msg)
}

func (u *UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id value", http.StatusBadRequest)
	}

	var user types.User

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	errDecode := decoder.Decode(&user)
	if errDecode != nil {
		log.Println("Error decoding JSON: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, msg := u.Db.UserRepo.UpdateUser(user, id)
	json.NewEncoder(w).Encode(response)
	json.NewEncoder(w).Encode(msg)
}

func (u *UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	userList, msg := u.Db.UserRepo.DeleteUser(id)
	json.NewEncoder(w).Encode(userList)
	json.NewEncoder(w).Encode(msg)
}
