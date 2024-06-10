package api

import (
	"github.com/gorilla/mux"
	"github.com/jtalev/gpg-staff-portal/api/handler"
)

func SetupRoutes(r *mux.Router) {

	// web
	r.HandleFunc("/", handler.LoginHandler)
	r.HandleFunc("/home", handler.HomeHandler)
	// r.HandleFunc("/clock-on", handler.ClockOnHandler)
	// r.HandleFunc("/leave", handler.LeaveHandler)

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
