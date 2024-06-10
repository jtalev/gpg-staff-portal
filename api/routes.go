package api

import "github.com/gorilla/mux"

func SetupRoutes(r *mux.Router) {
	// user
	r.HandleFunc("/user", nil).Methods("GET")
	r.HandleFunc("/user/{id}", nil).Methods("GET")
	r.HandleFunc("/user/{employeeId}", nil).Methods("GET")
	r.HandleFunc("/user/create", nil).Methods("POST")
	r.HandleFunc("/user/update", nil).Methods("PUT")
	r.HandleFunc("/user/delete", nil).Methods("DELETE")

	// timesheet

	// paycycle

	// leave request
}