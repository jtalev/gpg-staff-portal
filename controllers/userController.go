package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jtalev/gpg-staff-portal-api/persistence/mockRepos"
	"github.com/jtalev/gpg-staff-portal-api/entities"
)

type UserHandler struct {}

func (u *UserHandler) ServeHttp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Geelong Paint Group employee!"))
}

func (u *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	userList, statusMsg := persistence.GetAllUsers(false)

	json.NewEncoder(w).Encode(userList)
	json.NewEncoder(w).Encode(statusMsg)
}

func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	
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

	log.Println(user)

	response, statusMsg := persistence.CreateUser(user, false)

	json.NewEncoder(w).Encode(response)
	json.NewEncoder(w).Encode(statusMsg)
}

func (u *UserHandler) ReadUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	user, statusMsg := persistence.ReadUser(id)

	json.NewEncoder(w).Encode(user)
	json.NewEncoder(w).Encode(statusMsg)
}

func (u *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) 
	idStr := vars["id"]

	_, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id value", http.StatusBadRequest)
	}

	var user entities.User
	
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

	log.Println(user)
	response, statusMsg := persistence.UpdateUser(user, idStr)
	
	json.NewEncoder(w).Encode(response)
	json.NewEncoder(w).Encode(statusMsg)
}

func (u *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	userList, statusMsg := persistence.DeleteUser(id)
	json.NewEncoder(w).Encode(userList)
	json.NewEncoder(w).Encode(statusMsg)
}